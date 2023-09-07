package components

func RemoveDuplicates(wordsToCheck []string)[]string{
	UniqueMap :=make(map[string]bool)
	var uniqueValues []string
	for _, word := range wordsToCheck {
		if !UniqueMap[word]{
			UniqueMap[word] = true
			uniqueValues = append(uniqueValues, word)
		}
	}
	return uniqueValues
}