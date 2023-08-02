package components

func RemoveEmptyQuotes(wordlist []string)[]string{
	var newList []string;
	for i := 0; i < len(wordlist); i++ {
		if wordlist[i] != ""{
			newList = append(newList, wordlist[i])
		}
	}
	return newList
}