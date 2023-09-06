package main

import (
	"fmt"
	"net/http"

	handleapis "github.com/MarkVivian/wordGuesserHelper/handleApis"
)

var port int = 3001
func main() {
	http.HandleFunc("/randomWord", handleapis.RandomWord)
	http.HandleFunc("/searchWord", handleapis.SearchWord)

	fmt.Printf("Server listening to port %d \n ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil{
		fmt.Printf("Error while starting the server: %v \n", err)
	}
}