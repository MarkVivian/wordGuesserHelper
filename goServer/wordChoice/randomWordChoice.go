package wordchoice

import (
	"fmt"
	"math/rand"

	"github.com/MarkVivian/wordGuesserHelper/components"
	readdata "github.com/MarkVivian/wordGuesserHelper/readData"
)

var words []string = readdata.GetWords()

func RandomWordChooser(userlength int)string{ 
	var wordsMatchingLength = components.SortByLength(userlength, words)
	var randomPosition int = rand.Intn(len(wordsMatchingLength)) + 1
	var randomWord string = wordsMatchingLength[randomPosition]
	fmt.Printf("the word at position %d is %s", randomPosition, randomWord)
	return randomWord
}
