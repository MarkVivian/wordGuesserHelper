package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarkVivian/wordGuesserHelper/userchanges"
)

type clientInfo struct{
	Name string `json:"name"`
}

var port int = 3002
func main() {
	// starts the server.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from go server")
	})
	fmt.Printf("Server listening to port %d \n ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil{
		fmt.Printf("Error while starting the server: %v \n", err)
	}
}

func wordle(word string, length int, randomWord string){
	mapValues := userchanges.UserInput(word, length, randomWord)

	for key, value := range mapValues{
		if key != "hint"{
		fmt.Printf("the key is %s and the value is %v \n", key, value)
		}else{
			if value != ""{
				for keys, values := range value.(map[string]string) {
					fmt.Printf("the key is %s and the value is %v \n", keys, values)
				}
			}
		}
	}	
}