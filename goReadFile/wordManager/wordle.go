package wordmanager

import (
	"fmt"
	"math/rand"

	"github.com/MarkVivian/wordGuesserHelper/components"
	"github.com/MarkVivian/wordGuesserHelper/readfile"
)

var content []string = readfile.Words

func RandomWord(userLength int)string{
	randomValue := rand.Intn(len(components.ListSortByLength(userLength, content))) + 1
	randomword := components.ListSortByLength(userLength, content)[randomValue]
	fmt.Println(randomword)
	return randomword
}

func UserInput(word string, length int, randomWord string)components.WordleStruct{
	hint := make(map[string]string)
	wordList := components.ListSortByLength(length, content)
	for i := 0; i < len(wordList); i++ {
		if word == wordList[i] {
			if randomWord == word{
				return components.WordleStruct{
					Status : "correct",
					Reason : "correct word",
				}
			}else{
				for i := 0; i < len(randomWord); i++ {
					if word[i] == randomWord[i] {
						fmt.Printf("at %d word %s randomWord %s \n", i, string(word[i]), string(randomWord[i]))
						position := fmt.Sprintf("%d%s",i+1,components.GetNumberSuffix(i+1))
						hint[position] = string(word[i])
					}
				}
				return components.WordleStruct{
					Status : "wrong",
					Reason : "wrong word was given",
					Hint : hint,
 				}
			} 
		}
	}
	return components.WordleStruct{
		Status : "invalid",
		Reason : "word does not exist",
	}
}
