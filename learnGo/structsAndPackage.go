package main

import "fmt"

/*
		PACKAGE NAMES
for the main file it is given the package main.
the package name can be any value but is loved when you use the directory name EXCEPT FOR MAIN file.
if two or more files are in the same directory they are given the same package name.
		EXAMPLE
folder1
	|_ file1.go
	|_ file2.go

the package names should be similar so in this case we are going to use package folder1
*/

/*
		IMPORTING FILES.
mainFolder
	|____mainFile.go
		|__folder1
			|__ file1.go
			|__ file2.go

NB: BY VALUE I MEAN VARIABLE OR FUNCTION OR STRUCT .. ACTUALLY ANYTHING IMPORTABLE.
if you are importing a value from one file to another and they are in the same directory, file1 and file2, you just call the variable.
no export keyword neccesary.

if you are importing a value from another file, mainFile.go importing value from file1.go,
		1) the value that is being imported must start with capital letters.
		2) remember the module value in go.mod, we use that to import values in other directories.
				E.G.
lets say the module/go mod init value was path/example. then when importing a value, we are going to use:
	import (
		"fmt"
		"path/example/folderName" // since the folderName is the package of all the projects inside it.
	)

folderName.valueBeingImported // now you can access it using the folderName..REMEMBER IF IN DIFFERENT DIRECTORIES USE CAPITAL FIRST LETTER.

*/

// STRUCT
/*
similar to type in typescript but different😆
NB: IF THE VALUES ARE NOT GIVEN FOR THE DATA IN THE STRUCT THEY ARE GIVEN EITHER "" FOR STRING OR 0 FOR INT
*/
//CREATING STRUCT
type Anyname struct{
	// inside here we create the types similar to typescript interfaces.
	Name string
	Age uint8
	Gender string
	// we can also pass a struct as a datatype.
	work Bussiness
}

type Bussiness struct{
	salary uint16
	location string
}


			//USAGES 
func usingStructs() {
		// we can create an object of the struct in three ways.
		var name1 = Anyname{}
		name1.Age = 0
		name1.Name = "mark vivian"
		name1.Gender = "Male"
		name1.work.location = "mombasa"

		name2 := Anyname{
			Age: 10,
			Name: "Mary",
			Gender : "female",
			work: Bussiness{
				location: "nairobi",
				salary: 1000,
			},
		}
		fmt.Println(name2)

		type1 := Square{
			Lenght: 100,
			Width: 40,
		}

		fmt.Printf("the value is %v \n",type1.CalcArea())
		type1.changeValue(10)
		fmt.Printf("the value is %v \n", type1)
}

		//USING STRUCTS IN FUNCTIONS
type Square struct{
	Lenght uint8
	Width uint8
}
// this allows us to use the values of the object of the struct without having to pass them as variables.
func (S Square) CalcArea()uint8{
	return S.Lenght * S.Width
}

// we can also change the value using a pointer (*)
func (S *Square) changeValue(newValue int){
	S.Lenght = uint8(newValue)
}
