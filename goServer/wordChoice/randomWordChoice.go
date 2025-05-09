package wordchoice

import (
	"fmt"
	"math/rand"

	read_data "github.com/MarkVivian/wordGuesserHelper/read_data"
)

var words []string = read_data.GetWords()

func RandomWordChooser(userlength int) string {
	var wordsMatchingLength = SortByLength(userlength, words)
	var randomPosition int = rand.Intn(len(wordsMatchingLength)) + 1
	var randomWord string = wordsMatchingLength[randomPosition]
	fmt.Printf("the word at position %d is %s \n", randomPosition, randomWord)
	return randomWord
}

func SortByLength(length int, words []string) []string {
	var newList []string
	for i := 0; i <= (len(words) - 1); i++ {
		if len(words[i]) == length {
			newList = append(newList, words[i])
		}
	}
	return newList
}
