package main

import "fmt"

func main(){

    // define struct 
	type WelcomeUser struct{

		firstName string
		lastName string
		old uint
		email string
		
	}
		// example
	user := WelcomeUser {
			firstName :"kevin",
			lastName: "tido",
			old: 25,
			email: "kevintido@gmail.com",
		}
    
	// print 
	fmt.Printf(`

	Welcome %s %s.
	Your are %d years old. Your email is %s

	`, user.firstName, user.lastName, user.old, user.email)

}