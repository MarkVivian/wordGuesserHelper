package components

type RandomWordStruct struct{
	Length int `json:"length"`
	RandomWord string `json:"randomWord"`
}

type FindWordStruct struct{
	FirstLetter byte `json:"firstLetter"`
	LastLetter byte `json:"lastLetter"`
	Length int `json:"length"`
	LettersPresent []byte `json:"lettersPresent"`
	LettersNotPresent []byte `json:"lettersNotPresent"`
	WordList []string `json:"wordList"`
}