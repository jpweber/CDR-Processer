/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-01-29 14:57:23
 */

package main

import (
	"encoding/csv"
	// "encoding/json"
	"fmt"
	"os"
	// "strconv"
	"bluetone/cdrtest/CDRlib"
	// "reflect"
)

func main() {
	csvFile, err := os.Open("./1000309.ACT")
	// csvFile, err := os.Open("./data.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(csvData)

	// for _, value := range csvData[0] {
	// 	fmt.Println(value)
	// }

	cdrCollection := CDR.SplitTypes(csvData)

	stopRecords := make([]map[string]string, len(cdrCollection.Stops))
	for i, value := range cdrCollection.Stops {
		cdrStopData := CDR.FillCDRMap(CDR.CdrStopKeys(), value)
		stopRecords[i] = cdrStopData
	}

	fmt.Println("Done parsing file")

	//!debug
	// fmt.Println(stopRecords[0])
	json := CDR.JsonCdr(stopRecords[1])
	// fmt.Println(json)

}
