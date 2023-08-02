package components

func RemoveDuplicates(wordList []string)[]string{
	UniqueMap :=make(map[string]bool)
	var uniqueValues []string
	for _, word := range wordList {
		if !UniqueMap[word]{
			UniqueMap[word] = true
			uniqueValues = append(uniqueValues, word)
		}
	}
	return uniqueValues
}