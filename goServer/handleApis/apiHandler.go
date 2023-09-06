package handleapis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarkVivian/wordGuesserHelper/components"
	wordchoice "github.com/MarkVivian/wordGuesserHelper/wordChoice"
)

func RandomWord(w http.ResponseWriter,r *http.Request){
	var randomWordVariable components.RandomWordStruct
	err := json.NewDecoder(r.Body).Decode(&randomWordVariable)
	if err != nil{
		fmt.Printf("RandomWord: an error occured while reading the request for the random word line 17.... %v  \n", err.Error())
		return
	}
	randomWordVariable.RandomWord = wordchoice.RandomWordChooser(randomWordVariable.Length)
	randomWordJson, err := json.Marshal(randomWordVariable)
	if err != nil {
		fmt.Println("SendRandomWord : an error occured while converting the response to json form. \n" , err.Error())
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(randomWordJson)
}

func SearchWord(w http.ResponseWriter, r *http.Request){
	var searchWordsVariable components.FindWordStruct
	err := json.NewDecoder(r.Body).Decode(&searchWordsVariable)
	if err != nil{
		fmt.Printf("searchWord: an error occured while decoding the json to searchWordsVariable %v \n", err.Error())
		return
	}
	searchWordsVariable.WordList = wordchoice.FindRelatedWords(searchWordsVariable)
	fmt.Printf("the searchwordsvariable is %v \n", searchWordsVariable)
	/*
	searchWordsJson, err := json.Marshal(searchWordsVariable)
	if err != nil {
		fmt.Printf("searchRandomWords : an error occured while converting the searchWordsVariable to a json %v \n", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(searchWordsJson)
	*/
}