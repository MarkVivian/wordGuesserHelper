package components

func SortByLength(length int, words [] string) [] string{
	var newList []string
	for i := 0; i < len(words); i++ {
		if len(words[i]) == length {
			newList = append(newList, words[i])
		}
	}
	return newList
}