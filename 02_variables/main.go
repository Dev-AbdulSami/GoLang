package main

// numberOfUsers := 30000.15
const LoginToken string = "dsfjds" // public variable

// In Go, variable and constant visibility (public or private) is determined by the capitalization of the first letter of the variable name.

// If the name starts with an uppercase letter, the variable or constant is exported (i.e., public), meaning it can be accessed from other packages.
// If the name starts with a lowercase letter, the variable or constant is unexported (i.e., private), meaning it is only accessible within the same package.

// Since LoginToken starts with an uppercase letter, it is public and can be accessed from other packages

func main() {

	// var userName string = "Abdulsami"
	// fmt.Println(userName)
	// fmt.Printf("Variable is of type: %T \n", userName)

	// var isLoggedIn bool = true
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var smallValue int = 255
	// var smallValue uint8 = 256  //   ye error show kr rha he, because the range of uint8 is from 0 to 255
	// sb se acha he k int use krlen
	// fmt.Println(smallValue)
	// fmt.Printf("Variable is of type: %T \n", smallValue)

	// var smallFloatValue float64 = 255.65456456351351456
	// float32 . k bad k only 5 number show krta he
	// float64 . k bad k only 13 number show krta he
	// fmt.Println(smallFloatValue)
	// fmt.Printf("Variable is of type: %T \n", smallFloatValue)

	// default values and some aliases

	// var anotherVariables int
	// fmt.Println(anotherVariables) // ans 0 aega
	// fmt.Printf("Variable is of type: %T \n", anotherVariables)

	// var anotherVariables string
	// fmt.Println(anotherVariables) // empty aega
	// fmt.Printf("Variable is of type: %T \n", anotherVariables)

	// without datatype

	// var website = "youtube.com"
	// // agr koi datatype na den to jo value hogi us k according datatype set hojaegi, jese iski data type string he.
	// fmt.Println(website)
	// fmt.Printf("Variable is of type: %T \n", website)

	// no var style

	// Go programming language uses the := operator, which is known as the short variable declaration operator. It is used to declare and initialize variables in one step within a function's scope.

	// numberOfUsers := 30000.15 , ye syntax method k outside allow nhi he.
	// var numberOfUsers = 30000.15, method outaside k lie aese bnana hoga ya phir aese var numberOfUsers
	// float32 = 30000.15

	// numberOfUsers := 30000.15
	// fmt.Println(numberOfUsers)

	// fmt.Println(LoginToken)
	// fmt.Printf("Variable is of type: %T \n", LoginToken)
}
