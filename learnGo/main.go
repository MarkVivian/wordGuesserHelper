package main
import (
	"fmt"
)

func main() {
	fmt.Printf("this is go language \n");
	datatypes_and_variables();
	arrays_and_objects();
	usingStructs();
}

/*
				INITIALLIZATION
when you start a go project its good practice to initialize it by running:
	GO MOD INIT <<ANYTHING THAT WILL UNIQUELLY IDENTIFY THIS PROJECT>>
	GO MOD TIDY

this will create the go.mod file and that module name is going to be important when importing values.
*/

func datatypes_and_variables(){
	// all variables are declared with the data types.
	var name string = "mark vivian"
	var age int = 20
	var grade float32 = 2.45
	var change bool = false
	
	// you can also use short variable declaration to avoid the data type constrain.
	control := true;

	fmt.Printf("the name %s and age %d and change is %f and the grade is %t %v \n", name, age, grade, change, control)

	/*
	%s: Used for strings
	%d: Used for decimal integers (int, int8, int16, int32, int64)
	%f: Used for floating-point numbers (float32, float64)
	%t: Used for booleans
	%c: Used for individual characters (byte)
	%v: Used for any value (it will choose an appropriate format based on the type)
	%T: Used to print the type of the value
	*/

}

func arrays_and_objects(){
	// arrays are fixed size collections of the same type.
	var grades [5]int;
	grades[0] = 1;
	grades[2] = 4
	grades[1] = 3
	grades[4] = 9
	grades[3] = 10

	// arrays can also be abit loose with the size.
	var states = []string {"kenya", "usa", "america", "china"}
	fmt.Printf("the state values is %v \n", states)

	// this is like the type inference in typescript or struct in rust.
	type Names struct{
		Name string
		Age int
		gender string
	}
	// we pass the Name here 
	var names = [3]Names{
		{Name : "mark", Age: 12, gender: "wachira"},
		{Name : "tomas",Age : 10,gender : "anorld"},
		{Name : "apples",Age : 10,gender : "bread"},
	}
	fmt.Printf(names[0].Name)
}

func for_if_and_while_loops(){
	// for loops 
	var Numbers = [6]int {5,6,7,8,9,10}
	for index, value := range Numbers {
		fmt.Printf("Index: %d, value: %d \n", index, value)
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("the range is %d and the value is %d", i, Numbers[i])
	}

	if Numbers[3] > 5 {
		fmt.Print("the value is greater than 5")
	}else{
		fmt.Print("the value is less than 5")
	}
}