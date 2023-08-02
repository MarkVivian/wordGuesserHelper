package handlerequests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarkVivian/wordGuesserHelper/components"
	wordmanager "github.com/MarkVivian/wordGuesserHelper/wordManager"
)

func SendRandomWord(w http.ResponseWriter, r *http.Request){
	var randomWordVar components.RandomWordStruct
	err := json.NewDecoder(r.Body).Decode(&randomWordVar)
	if err != nil {
		fmt.Println("SendRandomWord : an error occured while decoding the request body.. ", err.Error())
		return
	}
	randomWordVar.RandomWord = wordmanager.RandomWord(randomWordVar.Length)
	randomWordJson, err := json.Marshal(randomWordVar)
	if err != nil {
		fmt.Println("SendRandomWord : an error occured while converting the response to json form." , err.Error())
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(randomWordJson)
}

func CheckIfCorrect(w http.ResponseWriter, r *http.Request){
	var whatToCheck components.CheckIfCorrectStruct
	err := json.NewDecoder(r.Body).Decode(&whatToCheck)
	if err != nil {
		fmt.Println("CheckIfCorrect : an error occured while decoding the request body.. ", err.Error())
		return
	}
	dataGotten := wordmanager.UserInput(whatToCheck.UserWord, whatToCheck.Length, whatToCheck.RandomWord)
	whatToCheck.Hint = dataGotten.Hint
	whatToCheck.Status = dataGotten.Status
	whatToCheck.Reason = dataGotten.Reason
	whatToCheckJson, err := json.Marshal(whatToCheck)
	if err != nil {
		fmt.Println("CheckIfCorrect : an error occured while converting the response to json form." , err.Error())
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(whatToCheckJson)
}

func FindWord(w http.ResponseWriter, r *http.Request){
	var wordToFind components.FindWordStruct
	err := json.NewDecoder(r.Body).Decode(&wordToFind)
	if err != nil {
		fmt.Println("FindWord : an error occured while decoding the request body.. ", err.Error())
		return
	}
	wordToFind.WordList = components.RemoveDuplicates(components.RemoveEmptyQuotes(wordmanager.WordFound(wordToFind)))
	wordToFindJson, err := json.Marshal(wordToFind)
	if err != nil {
		fmt.Println("FindWord : an error occured while converting the response to json form." , err.Error())
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(wordToFindJson)
}
