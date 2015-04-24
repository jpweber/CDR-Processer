/*
* @Author: jpweber
* @Date:   2015-04-23 20:50:06
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-23 21:12:33
 */

package FileHandling

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileList(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	fileNames := make([]string, 0)

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

func ArchiveFile(filename string) bool {
	fmt.Println(filename)
	err := os.Rename("ACT/"+filename, "ACT/archive/"+filename)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
