/*
* @Author: jpweber
* @Date:   2015-04-23 20:50:06
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-26 00:38:58
 */

package FileHandling

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func FileList(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	fileNames := make([]string, 0)

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, path.Join(dir, file.Name()))
		}
	}
	return fileNames
}

func ArchiveFile(filename string) bool {
	fmt.Println(filename)
	err := os.Rename(filename, path.Join(filepath.Dir(filename), "archive", path.Base(filename)))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
