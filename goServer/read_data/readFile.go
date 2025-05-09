package read_data

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetWords()[]string{ // this function will return a slice of strings.

	// load the path from the .env file.
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Error loading .env file: %v \n", err)
		os.Exit(1)
	}


	file, err := os.Open(os.Getenv("WORD_LIST_PATH")) // this will open the file in read mode.
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
