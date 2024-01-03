package main

import "fmt"

/*
			ERROR HANDLING
there some inbuild functions that have an error and the value as an output.
in this situations we use :
	valueStoredHere, err := function()
	if err != nil {
		fmt.Printf(err)
	}
*/
//lets make our own.
func calculate()(int, error){
	return 1,nil
}

func getCalc(){
	value, err := calculate()
	if err != nil {
		fmt.Printf("the error %v \n",err)
	}
	fmt.Println(value)
}