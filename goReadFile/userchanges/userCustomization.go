package userchanges

import (
	"fmt"
	"github.com/MarkVivian/wordGuesserHelper/readfile"
)


func Userinput()string{
	var userValue string;
	fmt.Println("guess the word : ")
	fmt.Scanln(&userValue)
	return userValue
}

func SortByLength()[]string{
	var length int = 0;
	var content []string = readfile.OpenFile()
	var newlist []string;
	fmt.Println("how long do you wish the word to be : (5 - 10)")
	fmt.Println("sort length is ", length)
	for i := 0; i < len(content); i++ {
		if len(content[i]) == length {
			newlist = append(newlist, content[i])
		}
	}
	/*
	fmt.Scan(&length)
	if length >= 5 && length <= 10 {

	}else{
		fmt.Printf("invalid value given.. try again: \n")
		SortByLength()
	}
	return newlist
	*/
}

func Chances()int{
	var userValue int;
	fmt.Println("how many chances do you want: (1 - 5) ")
	fmt.Scan(&userValue)
	if userValue > 5 || userValue < 1 {
		fmt.Printf("invalid value has been given.. try again: \n")
		Chances()
	}
	return userValue
}