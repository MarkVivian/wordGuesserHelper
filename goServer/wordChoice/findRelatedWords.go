package wordchoice

import (
	"fmt"

	"github.com/MarkVivian/wordGuesserHelper/components"
)

func FindRelatedWords(helper components.FindWordStruct) []string{
	var sortedValues []string = components.SortByLength(helper.Length, words)
	var newList []string = firstAndLastLetterChecker(helper.FirstLetter, helper.LastLetter, sortedValues)
	return checkRemainingLetters(newList, helper.LettersPresent)
}

func firstAndLastLetterChecker(firstLetter string, lastLetter string, wordsToCheck []string)[]string{
	var returningList []string
	var doneWithActivities []string
	if firstLetter == "" && lastLetter == "" {
		return wordsToCheck
	}else{
		if firstLetter != "" && lastLetter != ""{
			for _, firstwordValue := range wordsToCheck {
				if firstwordValue[0] == firstLetter[0]{
					// fmt.Printf("the word %v is similar to the first letter %v \n", firstwordValue, firstLetter)
					doneWithActivities = append(doneWithActivities, firstwordValue)
				}
			}
			for _, lastwordValue := range doneWithActivities {
				if lastwordValue[len(lastwordValue) - 1] == lastLetter[0]{
					// fmt.Printf("the word %s is similar to the last letter %s \n ", lastwordValue, lastLetter)
					returningList = append(returningList, lastwordValue)
				}
			}
		}else{
			if firstLetter != "" && lastLetter == ""{
				for _, firstwordValue := range wordsToCheck {
					if firstwordValue[0] == firstLetter[0]{
						// fmt.Printf("the word %v is similar to the first letter %v \n", firstwordValue, firstLetter)
						returningList = append(returningList, firstwordValue)
					}
				}
			}else if lastLetter != "" && firstLetter == ""{
				for _, lastwordValue := range wordsToCheck {
					if lastwordValue[len(lastwordValue)-1] == lastLetter[0]{
						// fmt.Printf("the word %v is similar to the last letter %v \n", lastwordValue, lastLetter)
						returningList = append(returningList, lastwordValue)
					}
				}
			}
		}
	}
	// fmt.Printf("the returning list is %v \n ", returningList)
	return returningList
}

func checkRemainingLetters(wordsToCheck []string, letterToUse []string)[]string{
	var remainingLetters []string
	for _, eachWord := range wordsToCheck{
		for _, eachLetter := range letterToUse {
			for i := 1; i < (len(eachWord)-1); i++ {
				if eachWord[i] == eachLetter[0]{
					remainingLetters = append(remainingLetters, eachWord)
				}
			}
		}
	}
	fmt.Printf("the none duplicate values are %v \n", components.RemoveDuplicates(remainingLetters))
	return components.RemoveDuplicates(remainingLetters)
}

/*
func FindRelatedWords(helper components.FindWordStruct)[]string{
	var filterStage0 []string = components.SortByLength(helper.Length, words)

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
}*/