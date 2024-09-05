package go_packages

import (
	"fmt"
)

type WelcomeUser struct{

		firstName string
		lastName string
		old uint
		email string
		
	}


func GetWelcomeUser(){

	// example
	user := WelcomeUser {
		firstName :"kevin",
		lastName: "tido",
        old: 25,
        email: "kevintido@gmail.com",
	}

	fmt.Println("Please Enter your firstName: ")
	fmt.Scan(&user.firstName)

	fmt.Println("Please Enter your lastName: ")
	fmt.Scan(&user.lastName)

    fmt.Println("How old you are: ")
	fmt.Scan(&user.old)

	fmt.Println("Please Enter your email: ")
	fmt.Scan(&user.email)


	fmt.Printf(`

	 Welcome %s %s.
	 Your are %d years old. Your email is %s

	 `, user.firstName, user.lastName, user.old, user.email)
 
}





