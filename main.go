/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-05-04 17:02:56
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
	// "encoding/json"
	"code.google.com/p/gcfg"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

//variables for displaying version information
const AppVersion = "0.9.2"

var buildNumber string

type Configuration struct {
	Required struct {
		FileDir string
		DbHost  string
		DbPort  int
		DbUser  string
		DBPass  string
		DSN     string
		FileExt string
	}
}

func saveRecord(wg *sync.WaitGroup, db sql.DB, records []map[string]string, recordType string, chunk *int, dbName string) {
	//make a buffered channel to hold all the records
	c := make(chan map[string]string, len(records)+1)
	for _, record := range records {
		// fill buffered channel with data
		c <- record
	}
	//close the buffered channel. This will hold all existing data
	//but will now take anymore new data. Works as a queue.
	close(c)

	for {

		tx, _ := db.Begin()

		var record map[string]string
		// pull records out of the buffered channel to insert to the database
		// pull out N (chunk) records then commit transaction.
		// Continue this until the channel is empty
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
			stmt, err := tx.Prepare("INSERT INTO " + dbName + "." + recordType + "(" + columnsString + ") VALUES(" + placeHoldersString + ")")
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

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "CARGO")
	if e == nil {
		log.SetOutput(logwriter)
	}
	log.SetFlags(0)
	log.Println("CARGO Starting")

	//load config from from config file
	// configFile, _ := os.Open("config.json")
	// decoder := json.NewDecoder(configFile)
	// configuration := Configuration{}
	// err := decoder.Decode(&configuration)
	var configuration Configuration
	err := gcfg.ReadFileInto(&configuration, "cargo.conf")
	if err != nil {
		fmt.Println("error:", err)
		log.Fatal("error", err)
	}

	//check for and create if needed the archive dir
	res := FileHandling.ArchivePrecheck(configuration.Required.FileDir)
	if res != true {
		os.Exit(1)
	}

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Then send a true bool to the wantToExit channel. The file archive process below
	// will then exit after it finishes any file work it is in the middle of.
	// This is a way to make clean exits if they are user intiated.
	signalChan := make(chan os.Signal, 1)
	wantToExit := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt, Waiting for archiving functions to complete...\n")
			log.Println("CARGO Received an interrupt, Waiting for archiving functions to complete")
			wantToExit <- true
		}
	}()

	for {

		var singleLoop bool

		//start timer for loop. This is just for performance information
		t0 := time.Now()

		var files []string
		if *cdrFileName != "" {
			fileToOpen := *cdrFileName
			files = append(files, fileToOpen)
			// if we are processing a single file to not loop around
			// aka do not run as a daemon
			singleLoop = true
		} else {
			files = FileHandling.FileList(configuration.Required.FileDir, configuration.Required.FileExt)
		}

		if len(files) == 0 {
			log.Println("No files waiting. Sleeping for 60 seconds")
			time.Sleep(time.Second * 60)
			continue
		} else {
			log.Printf("Processing %d Files\n", len(files))
		}

		// start db stuff
		db, err := sql.Open("mysql", configuration.Required.DSN)
		if err != nil {
			log.Fatal(err)
		}

		//test our connection to make sure the DB is reachable.
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		// I don't like this giant loop but its a simple way to start and test
		for _, file := range files {
			fmt.Println("starting loop")
			csvFile, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			reader := csv.NewReader(csvFile)
			reader.FieldsPerRecord = -1
			csvData, err := reader.ReadAll()

			if err != nil {
				log.Fatal(err)
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
			log.Println("Done parsing file")

			//get the dbname from the dsn in the config
			dsnParts := strings.Split(configuration.Required.DSN, "/")
			dbName := dsnParts[1]

			//Begin inserting CDR Data
			wg.Add(3)
			go saveRecord(&wg, *db, stopRecords, "stops", transactionChunk, dbName)
			go saveRecord(&wg, *db, attemptRecords, "attempts", transactionChunk, dbName)
			go saveRecord(&wg, *db, startRecords, "starts", transactionChunk, dbName)
			wg.Wait() //Wait for the concurrent routines to call 'done'

			//archive the raw files
			log.Println("Archiving " + file)
			res := FileHandling.ArchiveFile(file)
			time.Sleep(time.Second * 5)
			if res != true {
				log.Println("Error moving file")
				log.Println(file)
				os.Exit(1)
			} else {
				//if the archive succeeds check and see if a user
				//has tried to end the application
				//if they have then exit, otherewise keep going
				select {
				case exitStatus := <-wantToExit:
					if exitStatus {
						log.Println("Exiting Gracefully")
						os.Exit(0)
					}
				default:
					//do nothing
				}

			}

		}

		// if we are set to only loop once
		// return here to end the program
		if singleLoop {
			fmt.Println("Work Complete")
			return
		}

		fmt.Printf("%d Files took %f Seconds to process\n", len(files), time.Since(t0).Seconds())
		log.Printf("%d Files took %f Seconds to process\n", len(files), time.Since(t0).Seconds())
	} //end of main loop

}
