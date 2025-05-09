package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	handleapis "github.com/MarkVivian/wordGuesserHelper/handleApis"
)

func main() {
	erro := godotenv.Load(".env")

	if erro != nil {
		fmt.Printf("Error loading .env file: %v \n", erro)
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	http.HandleFunc("/randomWord", handleapis.RandomWord)
	http.HandleFunc("/searchWord", handleapis.SearchWord)

	fmt.Printf("Server listening to port %s \n ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil{
		fmt.Printf("Error while starting the server: %v \n", err)
	}
}