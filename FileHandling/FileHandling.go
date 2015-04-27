/*
* @Author: jpweber
* @Date:   2015-04-23 20:50:06
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-27 13:19:36
 */

package FileHandling

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func FileList(dir string, ext string) []string {

	files, err := filepath.Glob(dir + "/*." + ext)
	if err != nil {
		fmt.Println(err)
	}

	fileNames := make([]string, 0)

	for _, file := range files {
		fileNames = append(fileNames, file)
	}
	return fileNames
}

func ArchiveFile(filename string) bool {
	archivePath := path.Join(filepath.Dir(filename), "archive", path.Base(filename))
	err := os.Rename(filename, archivePath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	//now gzip the file
	//gzip the archived file
	res := CreateGZ(archivePath)
	if res != true {
		fmt.Println("Error GZipping file")
		fmt.Println(archivePath)
		os.Exit(1)
	}

	//delete the uncomressed file
	err = os.Remove(archivePath)
	if err != nil {
		fmt.Println("Error removing uncompressed file")
		fmt.Println(archivePath)
		os.Exit(1)
	}

	return true
}

func CreateGZ(filename string) bool {

	rawfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rawfile.Close()

	// calculate the buffer size for rawfile
	info, _ := rawfile.Stat()

	var size int64 = info.Size()
	rawbytes := make([]byte, size)

	// read rawfile content into buffer
	buffer := bufio.NewReader(rawfile)
	_, err = buffer.Read(rawbytes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	writer.Write(rawbytes)
	writer.Close()

	err = ioutil.WriteFile(filename+".gz", buf.Bytes(), info.Mode())
	// use 0666 to replace info.Mode() if you prefer

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s compressed to %s\n", filename, filename+".gz")

	return true
}

func ArchivePrecheck(cdrpath string) bool {
	//os.IsExist
	_, err := os.Stat(cdrpath + "/archive")
	if err != nil {
		fmt.Println("Does not exist")
		err = os.Mkdir(cdrpath+"/archive/", 0755)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			//after we create the dir make recursive call to thos function and run the checks again
			fmt.Println("calling check again")
			ArchivePrecheck(cdrpath)
		}
	} else {
		fmt.Println("File exists")
	}
	return true

}
