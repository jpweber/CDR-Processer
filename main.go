/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-20 21:11:21
 */

//parses CDR file in to key value map and then does something with it
// maybe database maybe not we'll see.

package main

import (
	// "bytes"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"ko/CDR"
	"os"
	"strings"
	"sync"
)

func saveRecord(wg *sync.WaitGroup, db sql.DB, records []map[string]string, recordType string) {
	for _, record := range records {
		//create strings for column names to insert
		//and string for all the values
		columns := make([]string, 0, len(record))
		stopsArgs := make([]interface{}, 0, len(record))
		placeHolders := make([]string, 0, len(record))
		for k, v := range record {
			columns = append(columns, k)
			stopsArgs = append(stopsArgs, v)
			placeHolders = append(placeHolders, "?")
		}

		columnsString := strings.Join(columns, ", ")
		// fmt.Println(stopsArgs)
		placeHoldersString := strings.Join(placeHolders, ", ")

		stmt, err := db.Prepare("INSERT INTO test." + recordType + "(" + columnsString + ") VALUES(" + placeHoldersString + ")")
		if err != nil {
			// log.Fatal(err)
			fmt.Println("prepare error")
			fmt.Println(err)
		}
		res, err := stmt.Exec(stopsArgs...)
		if err != nil {
			// log.Fatal(err)
			fmt.Println("Exec error")
			fmt.Println(err)
		}
		//debug info
		_, err = res.LastInsertId()
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}
		_, err = res.RowsAffected()
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}
		// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		// fmt.Printf("affected = %d\n", rowCnt)

	}
	wg.Done()
}

func main() {

	var buildNumber string
	const AppVersion = "0.1.0"
	versionPtr := flag.Bool("v", false, "Show Version Number")
	cdrFileName := flag.String("f", "", "Single file you which to process")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()
	if *versionPtr == true {
		fmt.Println(AppVersion + " Build " + buildNumber)
		os.Exit(0)
	}

	var fileToOpen string
	if *cdrFileName == "" {
		// csvFile, err := os.Open("./1000309.ACT")
		// csvFile, err := os.Open("./data.csv")
		fileToOpen = "./ACT/CHGOKBSBC01.20150301000000.100442B.ACT"
	} else {
		fileToOpen = *cdrFileName
	}

	//!debug
	fmt.Println(fileToOpen)
	// csvFile, err := os.Open("./1000309.ACT")
	// csvFile, err := os.Open("./data.csv")
	csvFile, err := os.Open(fileToOpen)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	// reader := csv.NewReader(bytes.NewReader(csvFile))
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cdrCollection := CDR.SplitTypes(csvData)
	if cdrCollection != nil {
		//if we've created a cdr collection we no longer need csvData
		//nil it out to release memory for GC
		csvData = nil
	}

	var wg sync.WaitGroup
	wg.Add(2) //We need to wait for Stops, attempts and starts to finish on this wait group

	//setup our containers for CDR data
	stopRecords := make([]map[string]string, len(cdrCollection.Stops))
	// create stop record map (dictionary to me)
	go func(wg *sync.WaitGroup) {
		// stopRecords := make([]map[string]string, len(cdrCollection.Stops))
		for i, value := range cdrCollection.Stops {
			cdrStopData := CDR.FillCDRMap(CDR.CdrStopKeys(), value) //normal
			//if we want to break out subfields
			cdrStopData = CDR.BreakOutSubFields(cdrStopData)
			stopRecords[i] = cdrStopData
		}
		wg.Done()
		fmt.Println("Stop Done")
	}(&wg)

	// create attempt record map (dictionary to me)
	go func(wg *sync.WaitGroup) {
		attemptRecords := make([]map[string]string, len(cdrCollection.Attempts))
		for i, value := range cdrCollection.Attempts {
			cdrAttemptData := CDR.FillCDRMap(CDR.CdrAttemptKeys(), value)
			attemptRecords[i] = cdrAttemptData
		}
		wg.Done()
		fmt.Println("Attempt Done")
	}(&wg)

	// create start record map (dictionary to me)
	// go func(wg *sync.WaitGroup) {
	// 	startRecords := make([]map[string]string, len(cdrCollection.Starts))
	// 	for i, value := range cdrCollection.Starts {
	// 		cdrStartData := CDR.FillCDRMap(CDR.CdrStartKeys(), value)
	// 		//if we want to break out subfields
	// 		cdrStartData = CDR.BreakOutSubFields(cdrStartData)
	// 		startRecords[i] = cdrStartData
	// 	}
	// 	wg.Done()
	// 	fmt.Println("Start Done")
	// }(&wg)

	wg.Wait() //Wait for the concurrent routines to call 'done'
	fmt.Println("Done parsing file")
	// fmt.Println(stopRecords[0])
	// fmt.Println(startRecords)

	// start db stuff
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	//test our connection to make sure the DB is reachable.
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	//Begin inserting CDR Data
	wg.Add(1) //will become three
	go saveRecord(&wg, *db, stopRecords, "stops")
	// go saveRecord(&wg, *db, attemptRecords, "attempts")
	// go saveRecord(&wg, *db, startRecords, "starts")
	wg.Wait() //Wait for the concurrent routines to call 'done'
	defer db.Close()

}
