package controller

import (
	"ExelProject/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SearchController(w http.ResponseWriter, request *http.Request) {
	fmt.Println("File Upload and Search endpoint")
	file, handler, err := request.FormFile("file")
	if err != nil {
		fmt.Println("Error in reading file")
	}

	fmt.Println("file size :", handler.Filename)
	defer file.Close()

	tempFile, err := ioutil.TempFile("data", "tempfile.xlsx")
	if err != nil {
		fmt.Println("Error in creating temp file")
	}
	defer tempFile.Close()

	value := request.Header.Get("key")
	if value == "" {
		w.Write([]byte("Please send the header key"))
		return
	}
	service.PostHandleAndSearch(file, value)
}
func del() {
	err := os.Remove("path")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File tempfile.xlsx successfully deleted")
}
func main() {
	if FileExists("path/to/file") {
		fmt.Println("Example file found.")
	} else {
		fmt.Println("Example file not found!")
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
