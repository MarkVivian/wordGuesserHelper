package main

import (
	"fmt"
	"math/rand"

	"github.com/MarkVivian/wordGuesserHelper/userchanges"
)

func main(){
	fmt.Println("LET THE GAME BEGIN")
	gameTime()
}

func gameTime(){
	curentList := userchanges.SortByLength()
	fmt.Println(curentList)
	randomValue := rand.Intn(len(curentList)) + 1
	var chances int = userchanges.Chances()
	var correctPosition []string;
	fmt.Printf("you have %d chances... ", chances)
	for i := 0; i < chances; i++ {
		inputGiven :=  userchanges.Userinput()
		if(len(inputGiven) == len(curentList[randomValue])){
			if(curentList[randomValue] == inputGiven){
				fmt.Println("congratulations you have guessed the correct value 🕺")
				return
			}else{
				for i := 0; i < len(inputGiven); i++ {
					if inputGiven[i] == curentList[randomValue][i] {
						correctPosition = append(correctPosition, string(inputGiven[i]))
					}
				}
				fmt.Printf("the position of \n %v \n was correct. try again. you have %d chances left: ", correctPosition, chances)
			}
		}else{
			fmt.Printf("the word %s is not %d letters. you have %d chances left ...", inputGiven, len(curentList[randomValue]), chances)
			i--
		}
	}
	fmt.Println("the word was ", )
}