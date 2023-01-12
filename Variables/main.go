package main

import "fmt"

const loginToken string = "AXsasasa" //Public

func main()  {
	fmt.Println(loginToken)
	fmt.Printf("Variable is of type: %T \n", loginToken)

	var username string = "Amol"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.36723642
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariablen int
	fmt.Println(anotherVariablen)
	fmt.Printf("Variable is of type: %T \n", anotherVariablen)

	// implicit type to declare variable
	var website = "learniongGo"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//No var style (:=)
	numberofUser := 5000
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of type: %T \n", numberofUser)
}