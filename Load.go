/*
* @Author: Jim Weber
* @Date:   2015-01-28 11:48:33
* @Last Modified by:   jpweber
* @Last Modified time: 2015-01-29 11:46:59
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

	cdrKeys := CDR.CdrStopKeys()
	// cdrKeys := make([]string, 0)
	// stopCDR := new(CDR.StopCDR)
	// s := reflect.ValueOf(stopCDR).Elem()
	// typeOfT := s.Type()
	// for i := 0; i < s.NumField(); i++ {
	// 	// f := s.Field(i)
	// 	// fmt.Printf("%d: %s \n", i, typeOfT.Field(i).Name)
	// 	cdrKeys = append(cdrKeys, typeOfT.Field(i).Name)
	// }

	// fmt.Println(cdrKeys)
	records := make([]map[string]string, 0)
	cdrData := make(map[string]string)
	for _, value := range csvData {
		// fmt.Println(i)
		// fmt.Println(value[2])
		// fmt.Println(reflect.ValueOf(value))
		cdrData = CDR.FillCDRMap(cdrKeys, value)
		records = append(records, cdrData)
	}

	// fmt.Println(cdrData)
	// fmt.Println(records)
	fmt.Println("Done parsing file")
	for key, value := range records[1] {
		fmt.Println(key, value, \n)
	}

	// CDR.JsonCdr(records[1])

}
