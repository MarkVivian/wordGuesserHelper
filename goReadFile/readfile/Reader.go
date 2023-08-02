package readfile

import (
	"bufio"
	"fmt"
	"os"
)

func openFile()[]string{
	file, err := os.Open("readfile/english3.txt")// the path should begin from the root directory.
	if err != nil {
		fmt.Println("an error occured while reading the file: ", err.Error())
		os.Exit(1)// this will stop the running of the program immediately
	}

	scanner := bufio.NewScanner(file)
	var valuesList []string;
	for scanner.Scan() {
		line := scanner.Text()
		valuesList = append(valuesList, line)
	}

	defer file.Close() // this will run when the entire function has run.

	error := scanner.Err()
	if error != nil{
		fmt.Printf("an error occured while reading the file %v \n", error.Error())
	}
	return valuesList
}

var Words []string = openFile()