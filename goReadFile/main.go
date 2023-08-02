package main

import (
	"fmt"
	"net/http"

	handlerequests "github.com/MarkVivian/wordGuesserHelper/handleRequests"
)

var port int = 3002

func main() {
	
	http.HandleFunc("/randomWord", handlerequests.SendRandomWord)
	http.HandleFunc("/checkIfCorrect", handlerequests.CheckIfCorrect)
	http.HandleFunc("/findWord", handlerequests.FindWord)

	fmt.Printf("Server listening to port %d \n ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil{
		fmt.Printf("Error while starting the server: %v \n", err)
	}
}

