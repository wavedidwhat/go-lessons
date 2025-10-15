package main

import (
	"fmt"
	"github.com/wavedidwhat/go-lessons/week_1/sample/what-type/constants"
	"github.com/wavedidwhat/go-lessons/week_1/sample/what-type/types" // Correct import path
	"github.com/wavedidwhat/go-lessons/week_1/sample/what-type/variables"
)

func main() {
	// fmt.Println(constants.userName)
	fmt.Println(constants.Username)
	fmt.Println(constants.Age)
	fmt.Println(constants.Publicusername)
	// Using types package
	var user types.User
	user.ID = types.UserID(1)
	user.UserName = types.UserName("Alice")
	user.Body = types.UserBody(65.5)
	user.Active = types.UserActive(true)

	fmt.Println(user.DisplayInfo())
	fmt.Println(user.BalanceString())

	// Using variables package
	fmt.Println(variables.Publicusername)
	// fmt.Println(variables.privateUsername) // This will cause a compilation error
	fmt.Println(variables.Name)
	// fmt.Println(variables.age) // This will cause a compilation error
	// you can dynamically change the value of a variable
	variables.Name = "newWave"
	fmt.Println(variables.Name)
	// using the short variable declaration inside a function
	age := 24
	fmt.Println(age)
}