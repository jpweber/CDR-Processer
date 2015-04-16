/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-16 17:39:23
 */

//parses CDR file in to key value map and then publishes to rabbitMQ

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"ko/CDR"
	"os"
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
	wg.Add(3) //We need to wait for Stops, attempts and starts to finish on this wait group

	// create stop record map (dictionary to me)
	go func(wg *sync.WaitGroup) {
		stopRecords := make([]map[string]string, len(cdrCollection.Stops))
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
	go func(wg *sync.WaitGroup) {
		startRecords := make([]map[string]string, len(cdrCollection.Starts))
		for i, value := range cdrCollection.Starts {
			cdrStartData := CDR.FillCDRMap(CDR.CdrStartKeys(), value)
			//if we want to break out subfields
			cdrStartData = CDR.BreakOutSubFields(cdrStartData)
			startRecords[i] = cdrStartData
		}
		wg.Done()
		fmt.Println("Start Done")
	}(&wg)

	wg.Wait() //Wait for the concurrent routines to call 'done'
	fmt.Println("Done parsing file")
	// fmt.Println(stopRecords)
	// fmt.Println(startRecords)
}
