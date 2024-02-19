package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var newFile *os.File
	_ = newFile
	var err error

	var userFilename string

	fmt.Println("hello welvome to file creating app made by manoj")

	fmt.Printf("can you please enter the name of your file\n:")
	fmt.Scanln(&userFilename)
	outputtext := checkUserFileName(userFilename)
	fmt.Println(outputtext)

	if outputtext == "file can be created" {

		var fullFileName = userFilename + ".txt"
		newFile, err = os.Create(fullFileName)
		if err != nil {
			log.Fatal(err)
		}
		var fileinfo os.FileInfo
		fileinfo, err := os.Stat(fullFileName)
		if err != nil {
			fmt.Println("sorry we are facing the data while getting the user ")

		} else {
			fmt.Println("file Info:")
			fmt.Println("FileName:", fileinfo.Name())

			fmt.Println("FileSize: ", fileinfo.Size())

			fmt.Println("Modified time:", fileinfo.ModTime())
			fmt.Println("is Directory: ", fileinfo.IsDir())
			newFile.Close()
		}

	} else {
		fmt.Println("exited program.")
	}

}

func checkUserFileName(userfilename string) string {
	var sendString string
	if userfilename != " " && userfilename != "" {
		if len(userfilename) > 10 {
			sendString = "length to long can't crete file."
			return sendString
		}
		sendString = "file can be created"
		return sendString
	} else {
		fmt.Println("sorry user file can't be created")
		sendString = "Invalid can't crete file."
		return sendString
	}
}
