package components




func ListSortByLength(userLength int, content []string) []string {
	var newlist []string
	for i := 0; i < len(content); i++ {
		if len(content[i]) == userLength {
			newlist = append(newlist, content[i])
		}
	}
	return newlist
}