package userchanges

import (
	"fmt"
	"math/rand"

	"github.com/MarkVivian/wordGuesserHelper/components"
	"github.com/MarkVivian/wordGuesserHelper/readfile"
)

var content []string = readfile.Words

func listSortByLength(userLength int) []string {
	var newlist []string
	for i := 0; i < len(content); i++ {
		if len(content[i]) == userLength {
			newlist = append(newlist, content[i])
		}
	}
	return newlist
}

func RandomWord(userLength int)string{
	randomValue := rand.Intn(len(listSortByLength(userLength))) + 1
	randomword := listSortByLength(userLength)[randomValue]
	fmt.Println(randomword)
	return randomword
}

func UserInput(word string, length int, randomWord string) map[string]interface{}{
	hint := make(map[string]string)
	wordList := listSortByLength(length)
	for i := 0; i < len(wordList); i++ {
		if word == wordList[i] {
			if randomWord == word{
				return map[string] interface{}{
					"status" : "correct",
					"reason" : "correct word",
					"hint" : "",
				}
			}else{
				for i := 0; i < len(randomWord); i++ {
					if word[i] == randomWord[i] {
						fmt.Printf("at %d word %s randomWord %s \n", i, string(word[i]), string(randomWord[i]))
						position := fmt.Sprintf("%d%s",i,components.GetNumberSuffix(i+1))
						hint[position] = string(word[i])
					}
				}
				return map[string]interface{}{
					"status" : "wrong",
					"reason" : "wrong word was given",
					"hint" : hint,
 				}
			}
		}
	}
	return map[string]interface{}{
		"status" : "invalid",
		"reason" : "word does not exist",
		"hint" : "",
	}
}