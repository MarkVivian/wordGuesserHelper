package wordchoice

import (
	"fmt"
	// "strconv"
	"github.com/MarkVivian/wordGuesserHelper/components"
	"github.com/MarkVivian/wordGuesserHelper/read_data"
)

var word_list []string = read_data.GetWords()

func FindRelatedWords(word_formatter components.FindWordStruct) []string {
	// convert the word_formatter.Length to int 
	// Length_converted, err := strconv.Atoi(word_formatter.Length)
	// if err != nil {
	// 	fmt.Printf("RandomWord: could not convert Length to int: %v\n", err.Error())
	// 	return nil
	// }
	var sortedValues []string = SortByLength(word_formatter.Length, word_list)
	var newList []string = firstAndLastLetterChecker(word_formatter.FirstLetter, word_formatter.LastLetter, sortedValues)
	return checkRemainingLetters(newList, word_formatter.LettersPresent, word_formatter.LettersNotPresent, sortedValues)
}

func firstAndLastLetterChecker(firstLetter string, lastLetter string, wordsToCheck []string) []string {
	var returningList []string
	var doneWithActivities []string
	// check if the first letter and last letter are empty
	// if they are empty, then return the wordsToCheck
	if firstLetter == "" && lastLetter == "" {
		fmt.Printf("the first letter is empty and the last letter is empty \n")
		return wordsToCheck
	} else {
		if firstLetter != "" {
			for _, firstwordValue := range wordsToCheck {
				// check if the first letter of the word is equal to the first letter of the word formatter
				if firstwordValue[0] == firstLetter[0] {
					// fmt.Printf("the word %s is similar to the first letter %s \n", firstwordValue, string(firstLetter))
					doneWithActivities = append(doneWithActivities, firstwordValue)
				}
			}
		} else {
			doneWithActivities = wordsToCheck
		}

		if lastLetter != "" {
			for _, lastwordValue := range doneWithActivities {
				if lastwordValue[len(lastwordValue)-1] == lastLetter[0] {
					// fmt.Printf("the word %s is similar to the last letter %s \n ", lastwordValue, string(lastLetter))
					returningList = append(returningList, lastwordValue)
				}
			}
		} else {
			returningList = doneWithActivities
		}
	}
	fmt.Printf("the returning list is %v \n ", returningList)
	return returningList
}

func checkRemainingLetters(wordsToCheck []string, letterToUse []string, letterToAvoid []string, empty_check []string) []string {
	var remainingLetters []string
	var finalList []string
	// check if letterToAvoid is empty
	if len(letterToAvoid) == 0 && len(letterToUse) == 0 {
		fmt.Printf("the letter to avoid is empty and the letter to use is empty \n")
		finalList = wordsToCheck
	} else {
		// check if letterToAvoid is not empty
		if len(letterToAvoid) != 0 {
			// check for each word in the wordsToCheck if it contains any of the letters in letterToAvoid
			// if it does not contain any of the letters in letterToAvoid, then add it to the remainingLetters
			for _, eachWord := range wordsToCheck {
				avoid_state := false
				for _, eachLetter := range letterToAvoid {
					for i := 0; i <= (len(eachWord) - 1); i++ {
						if eachWord[i] == eachLetter[0] {
							avoid_state = true
						}
					}
					if avoid_state {
						break
					}
				}
				if ! avoid_state {
					remainingLetters = append(remainingLetters, eachWord)
				}
			}
		} else {
			remainingLetters = wordsToCheck
			fmt.Printf("the letter to avoid is empty \n and the remaining letters are %v \n", remainingLetters)
		}
		// check if letterToUse is not empty
		if len(letterToUse) != 0 {
			//check for each word in the wordsToCheck if it contains any of the letters in letterToUse
			// if it contain all of the letters in letterToUse, then add it to the remainingLetters
			for _, eachWord := range remainingLetters {
				// state will ensure that all the letters in letterToUse are present in the word
				letter_state := false
				for _, eachLetter := range letterToUse {
					fmt.Printf("the letter to use is %s \n", string(eachLetter))
					for i := 0; i <= (len(eachWord) - 1); i++ {
						if eachWord[i] == eachLetter[0] {
							fmt.Printf("the word %s does contain the letter %s \n", string(eachWord[i]), string(eachLetter))
							// if the letter is found, state is set to true then it break to go to the next letter.
							letter_state = true
							break
						}
						fmt.Printf("the word %s does not contain the letter %s \n", string(eachWord[i]), string(eachLetter))
					}
					// if the letter_state is false, then it means that not all the letters in letterToUse are present in the word
					if letter_state {
						break
					}
				}
				// if the letter_state is true, then it means that all the letters in letterToUse are present in the word
				if letter_state {
					finalList = append(finalList, eachWord)
				}
			}

		} else {
			fmt.Printf("the letter to use is empty \n")
			finalList = remainingLetters
		}
	}
	fmt.Printf("the final list is are %v \n", finalList)
	return finalList
}
