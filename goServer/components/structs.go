package components

type RandomWordStruct struct{
	Length int `json:"length"`
	RandomWord string `json:"randomWord"`
}

type FindWordStruct struct{
	FirstLetter string `json:"firstLetter"`
	LastLetter string `json:"lastLetter"`
	Length int `json:"length"`
	LettersPresent []string `json:"lettersPresent"`
	LettersNotPresent []string `json:"lettersNotPresent"`
	WordList []string `json:"wordList"`
}