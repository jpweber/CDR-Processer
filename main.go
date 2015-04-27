/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-27 17:40:22
 */

//parses CDR file in to key value map and then does something with it
// maybe database maybe not we'll see.

package main

import (
	// "bytes"
	"CDR-Processer/CDR"
	"CDR-Processer/FileHandling"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"sync"
	"time"
)

//variables for displaying version information
const AppVersion = "0.9.0"

var buildNumber string

type Configuration struct {
	FileDir string
	DbHost  string
	DbPort  int
	DbUser  string
	DBPass  string
	DSN     string
	FileExt string
}

func saveRecord(wg *sync.WaitGroup, db sql.DB, records []map[string]string, recordType string, chunk *int) {
	//make a buffered channel to hold all the records
	c := make(chan map[string]string, len(records)+1)
	for _, record := range records {
		// fill buffered channel with data
		c <- record
	}
	close(c)

	for {

		tx, _ := db.Begin()

		var record map[string]string
		for i := 0; i < *chunk; i++ {
			record = <-c

			//if the chanel is empty it will start returning zeros
			if len(record) == 0 {
				tx.Commit()
				wg.Done()
				return
			}

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
			// fmt.Println(columnsString)
			placeHoldersString := strings.Join(placeHolders, ", ")

			//create the prepared statment
			stmt, err := tx.Prepare("INSERT INTO test." + recordType + "(" + columnsString + ") VALUES(" + placeHoldersString + ")")
			if err != nil {
				// log.Fatal(err)
				fmt.Println("prepare error")
				fmt.Println(err)
			}

			//execute the prepared statment
			res, err := stmt.Exec(stopsArgs...)
			if err != nil {
				// log.Fatal(err)
				fmt.Println("Exec error")
				fmt.Println(err)
			}

			//close the statment
			stmt.Close()

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

		}
		//!debug
		// fmt.Println(" commiting")
		tx.Commit()
	}

	// wg.Done()
}

func main() {

	versionPtr := flag.Bool("v", false, "Show Version Number")
	cdrFileName := flag.String("f", "", "Single file you wish to process")
	transactionChunk := flag.Int("t", 50, "Number of records to insert in a single transaction. experminting with this number can provide better perfomance sometimes. 500 and 1000 have been tested on a laptop")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()
	if *versionPtr == true {
		fmt.Printf("%s %s %s", AppVersion, "Build", buildNumber)
		os.Exit(0)
	}

	//load config from from config file
	configFile, _ := os.Open("config.json")
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	//check for and create if needed the archive dir
	res := FileHandling.ArchivePrecheck(configuration.FileDir)
	if res != true {
		os.Exit(1)
	}

	for {

		var singleLoop bool
		// var fileToOpen string
		var files []string
		if *cdrFileName != "" {
			fileToOpen := *cdrFileName
			files = append(files, fileToOpen)
			// if we are processing a single file to not loop around
			// aka do not run as a daemon
			singleLoop = true
		} else {
			files = FileHandling.FileList(configuration.FileDir, configuration.FileExt)
		}

		if len(files) == 0 {
			fmt.Println("Not files waiting. Sleeping for 60 seconds")
			time.Sleep(time.Second * 60)
			continue
		}

		// start db stuff
		db, err := sql.Open("mysql", configuration.DSN)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}

		//test our connection to make sure the DB is reachable.
		err = db.Ping()
		if err != nil {
			fmt.Println(err)
		}

		defer db.Close()

		// I don't like this giant loop but its a simple way to start and test
		for _, file := range files {
			csvFile, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
			}

			reader := csv.NewReader(csvFile)
			reader.FieldsPerRecord = -1
			csvData, err := reader.ReadAll()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			csvFile.Close()

			cdrCollection := CDR.SplitTypes(csvData)
			if cdrCollection != nil {
				//if we've created a cdr collection we no longer need csvData
				//nil it out to release memory for GC
				csvData = nil
			}

			var wg sync.WaitGroup
			wg.Add(3) //We need to wait for Stops, attempts and starts to finish on this wait group
			//If more go funcs are added the integer here needs increased to match

			//setup our containers for CDR data
			stopRecords := make([]map[string]string, len(cdrCollection.Stops))
			attemptRecords := make([]map[string]string, len(cdrCollection.Attempts))
			startRecords := make([]map[string]string, len(cdrCollection.Starts))

			//Populate the containers of CDR Data
			go func() {
				stopRecords = (CDR.CreateRecordMap(&wg, cdrCollection.Stops, "stops"))
			}()
			go func() {
				attemptRecords = (CDR.CreateRecordMap(&wg, cdrCollection.Attempts, "attempts"))
			}()
			go func() {
				startRecords = (CDR.CreateRecordMap(&wg, cdrCollection.Starts, "starts"))
			}()

			wg.Wait() //Wait for the concurrent routines to call 'done'
			fmt.Println("Done parsing file")

			//Begin inserting CDR Data
			wg.Add(3)
			go saveRecord(&wg, *db, stopRecords, "stops", transactionChunk)
			go saveRecord(&wg, *db, attemptRecords, "attempts", transactionChunk)
			go saveRecord(&wg, *db, startRecords, "starts", transactionChunk)
			wg.Wait() //Wait for the concurrent routines to call 'done'

			//archive the raw files
			res := FileHandling.ArchiveFile(file)
			if res != true {
				fmt.Println("Error moving file")
				fmt.Println(file)
				os.Exit(1)
			}
		}

		// if we are set to only loop once
		// return here to end the program
		if singleLoop {
			fmt.Println("Work Complete")
			return
		}
	} //end of main loop

}
