package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// setupDir sans dir for "files". It is used as the destination for data to encrypt.
// It is created if needed and the program returs.
func setupDir() {
	dirName := "files"
	path := "./" + dirName
	fmt.Printf("\nCheking directory ... ")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Printf("Folder %q successfully created.\n", dirName)
		fmt.Printf("\nPlease copy data into %q and restart the program.\n", dirName)
		exit()
	} else {
		fmt.Printf("Folder %q already exsists.\n", dirName)
	}
}

// mkDir scans dir for desired folder and creates it if necessary
func mkDir(dirName string) {

	path := "./" + dirName
	fmt.Printf("\nCheking directory ... ")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Printf("Folder %q successfully created.\n", dirName)
	} else {
		fmt.Printf("Folder %q already exsists.\n", dirName)
	}
}

// getEcnypted retruns the file information in dir: "encrypted"
func getEncrypted() (encs []string) {

	files, err := ioutil.ReadDir("./encrypted")
	if err != nil {
		fmt.Println("Reading encrypted file information failed! Program is terminated! Error: ", err)
		return
	}

	for _, file := range files {
		n := file.Name()
		encs = append(encs, n)
	}
	return
}

// exit() returns from the function
func exit() {
	fmt.Printf("\nPress %q to Exit.", "Enter")
	var input string
	fmt.Scanln(&input)
	return
}