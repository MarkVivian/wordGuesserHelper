package handleapis

import (
	"bytes"
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
	searchWordsJson, err := json.Marshal(sendToPythonResponseReciever(searchWordsVariable))
	if err != nil {
		fmt.Printf("searchRandomWords : an error occured while converting the searchWordsVariable to a json %v \n", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(searchWordsJson)
}

func sendToPythonResponseReciever(searchWordData components.FindWordStruct)components.FindWordStruct{
	var url string = "http://localhost:3002/filterWords"
	jsonData, err := json.Marshal(searchWordData)
	if err != nil{
		fmt.Printf("sendToPythonError: could not marshal the searchWordsLsit to jsonData %v \n", err.Error())
		return searchWordData
	}
	// Create an io.Reader from the JSON data
    jsonReader := bytes.NewReader(jsonData)
	response, err := http.Post(url, "application/json", jsonReader)
	if err != nil{
		fmt.Printf("sendToPythonError: could not post the data to the python server: %v \n ", err.Error())
		return searchWordData
	} 
	defer response.Body.Close()
	
	var responseUsingStruct components.FindWordStruct
	if err := json.NewDecoder(response.Body).Decode(&responseUsingStruct); err != nil{
		fmt.Printf("sendToPythonError: could not decode the response from python to the struct %v \n", err.Error())
		return searchWordData
	}
	return responseUsingStruct
}
