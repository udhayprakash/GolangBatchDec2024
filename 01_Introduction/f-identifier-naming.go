package main

import "fmt"

/*
variable, Identifier, user-defined variable

	RULES:
	-----
		1. first letter should be alphabet or _
		2. second letter onwards a-z, A-Z, _, 0-9
*/

func main() {

	// Identifiers :- user-defined variables
	x := "Jaleela"
	fmt.Println("Student name is ", x)

	// NOTE: code will be better readable if meaningful names are given for variables

	name := "Jaleela"
	fmt.Println("Student name is ", name)

	studentName := "Jaleela"
	fmt.Println("Student name is ", studentName)

	student2 := "Jaleela"
	fmt.Println("Student name is ", student2)

	// 2ndName := "Jaleela" // can't start with number
	// $name = "Jaleela"     // $ can't be used in identifier name

	_name := "Jaleela"
	fmt.Println("Student's name is ", _name)

	__name := "Jaleela"
	fmt.Println("Student's name is ", __name)

	__name__ := "Jaleela"
	fmt.Println("Student's name is ", __name__)


}


/*
NOTE:
	1. camelCasing is preferred for the identifiers
	2. Only variables, functions or types whose names begin with a capital
	letter are considered as exported: accessible from packages outside
	the current package.
	3. Names starting with a capital letter are public, names with a
	lowercase letter are private.
*/