package wordchoice

import (
	"github.com/MarkVivian/wordGuesserHelper/components"
)

var filterStage0, filterStage1, filterLastStage []string

func FindRelatedWords(helper components.FindWordStruct)[]string{
	filterStage0 = components.SortByLength(helper.Length, words)

	if helper.FirstLetter == "" && helper.LastLetter == ""{
		filterStage1 = words;
	}else{
		if helper.FirstLetter == ""{
			filterStage1 = checkFirstAndLastLetters("0", helper.LastLetter)
		}else{
			filterStage1 = checkFirstAndLastLetters(helper.FirstLetter, "0")
		}
	}

	filterLastStage = filterInnerWords(filterStage1, helper.LettersPresent)
	return filterLastStage
}

func checkFirstAndLastLetters(firstLetter string, lastLetter string)[]string{
	var firstLetterWordList, lastLetterWordList []string;
	if firstLetter != "0"{
		for _, word := range filterStage0 {
			if word[0] == firstLetter[0]{
				firstLetterWordList = append(firstLetterWordList, word)
			}
		}
	}else{
		firstLetterWordList = words;
	}
	if firstLetter != "0" {
		for _, word := range firstLetterWordList{
			if word[len(firstLetterWordList)-1] == lastLetter[0]{
				lastLetterWordList = append(lastLetterWordList, word)
			}
		}
	}else{
		return firstLetterWordList
	}
	return lastLetterWordList
}

func filterInnerWords(lastStage []string, checkerWords []string)[]string{
	var lastList []string;
	for _, word := range lastStage {
		
		// for i := 0; i < len(checkerWords); i++ {
		// 	for j := 0; j < len(word); j++ {
		// 		if word[j] == checkerWords[i][0]{
		// 			lastList = append(lastList, word)
		// 		}
		// 	}
		// }			
	}
	return lastList
}