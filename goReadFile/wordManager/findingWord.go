package wordmanager

import (
	"strings"

	"github.com/MarkVivian/wordGuesserHelper/components"
	"github.com/MarkVivian/wordGuesserHelper/readfile"
)

func WordFound(info components.FindWordStruct)[]string{
	var LastWordlist []string;
	var wordList []string = components.ListSortByLength(info.Length, readfile.Words)
	for i := 0; i < len(wordList); i++ {
		LastWordlist = append(LastWordlist, checkWithAvailableWords(info.LettersPresent, findByCharacter(rune(info.FirstLetter[0]), rune(info.LastLetter[0]), wordList[i])))
	}
	return LastWordlist
}

func findByCharacter(firstLetter rune, lastLetter rune, wordListValue string)string{
	var wordsThatMatch string; 
	if firstLetter != 0 {
		if firstLetter == rune(wordListValue[0]){
			wordsThatMatch = wordListValue
		}
	}
	if lastLetter != 0 {
		if lastLetter != rune(wordListValue[len(wordListValue) - 1]) {
			wordsThatMatch = ""
			return wordsThatMatch
		}
	}
	return wordsThatMatch
}

func checkWithAvailableWords(availableLetters []string, wordThatMatched string)string{
	var wordsThatFit string;
	if wordThatMatched != ""{
		for i := 0; i < len(availableLetters); i++ {
			
		}
		for _, letter := range availableLetters {
			if strings.ContainsRune(wordThatMatched[1:len(wordThatMatched)-2], rune(letter[0])) {
				wordsThatFit = wordThatMatched
			}
		}
	}
	return wordsThatFit
}