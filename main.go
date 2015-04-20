/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-20 00:28:50
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
	wg.Add(1) //We need to wait for Stops, attempts and starts to finish on this wait group

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
	// go func(wg *sync.WaitGroup) {
	// 	attemptRecords := make([]map[string]string, len(cdrCollection.Attempts))
	// 	for i, value := range cdrCollection.Attempts {
	// 		cdrAttemptData := CDR.FillCDRMap(CDR.CdrAttemptKeys(), value)
	// 		attemptRecords[i] = cdrAttemptData
	// 	}
	// 	wg.Done()
	// 	fmt.Println("Attempt Done")
	// }(&wg)

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
	fmt.Println(len(stopRecords[0]))
	// fmt.Println(startRecords)

	// stopsColumns := CDR.KeysString(stopRecords[0])
	// fmt.Println(stopsColumns)
	// stopsValues := CDR.ValuesString(stopRecords[0])
	// fmt.Println(stopsValues)
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
	//create strings for column names to insert
	//and string for all the values
	columns := make([]string, 0, len(stopRecords[0]))
	stopsValues := make([]string, 0, len(stopRecords[0]))
	placeHolders := make([]string, 0, len(stopRecords[0]))
	for k, v := range stopRecords[0] {
		columns = append(columns, k)
		stopsValues = append(stopsValues, v)
		placeHolders = append(placeHolders, "?")
	}

	columnsString := strings.Join(columns, ", ")
	stopsValuesString := strings.Join(stopsValues, "', '")
	placeHoldersString := strings.Join(placeHolders, ", ")

	stmt, err := db.Prepare("INSERT INTO test.stops(" + columnsString + ") VALUES(" + placeHoldersString + ")")
	if err != nil {
		// log.Fatal(err)
		fmt.Println("prepare error")
		fmt.Println(err)
	}
	res, err := stmt.Exec(stopsValuesString)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	defer db.Close()

}
