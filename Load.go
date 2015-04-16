/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-16 11:20:56
 */

//parses CDR file in to key value map and then publishes to rabbitMQ

package main

import (
	"encoding/csv"
	// "encoding/json"
	"fmt"
	"os"
	// "strconv"
	"ko/CDRlib"
	// "bluetone/cdrtest/Queuelib"
	// "reflect"
	//"sync"
)

func main() {

	//standard methods
	// csvFile, err := os.Open("./1000309.ACT")
	// csvFile, err := os.Open("./data.csv")
	csvFile, err := os.Open("./Stop.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
	// standard methods

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

	// create stop record map (dictionary to me)
	stopRecords := make([]map[string]string, len(cdrCollection.Stops))
	for i, value := range cdrCollection.Stops {
		cdrStopData := CDR.FillCDRMap(CDR.CdrStopKeys(), value) //normal
		//if we want to break out subfields
		cdrStopData = CDR.BreakOutSubFields(cdrStopData)
		stopRecords[i] = cdrStopData
	}

	// create attempt record map (dictionary to me)
	attemptRecords := make([]map[string]string, len(cdrCollection.Attempts))
	for i, value := range cdrCollection.Attempts {
		cdrAttemptData := CDR.FillCDRMap(CDR.CdrAttemptKeys(), value)
		attemptRecords[i] = cdrAttemptData
	}

	fmt.Println("Done parsing file")
	// go publish(c, stopRecords, &wg)
	// go publish(c, attemptRecords, &wg)
	fmt.Println(stopRecords)
}
