package components

type RandomWordStruct struct{
	Length int `json:"length"`
	RandomWord string `json:"randomWord"`
}

type WordleStruct struct{
	Status string `json:"status"`
	Hint map[string]string `json:"hint"`
	Reason string `json:"reason"`
}

type CheckIfCorrectStruct struct{
	RandomWordStruct
	WordleStruct
	UserWord string `json:"userWord"`
}

type FindWordStruct struct{
	FirstLetter string `json:"firstLetter"`
	LastLetter string `json:"lastLetter"`
	Length int `json:"length"`
	LettersPresent []string `json:"lettersPresent"`
	WordList []string `json:"wordList"`
}