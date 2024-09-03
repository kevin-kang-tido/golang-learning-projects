package main

import (
	"fmt" // this package use for O/I // "time" // playing with time
	"go_project/go_packages"
	"go_project/integers"
)

// write a function to welcome program
func GreeUsers(){

    fmt.Println("=======================================")
	fmt.Println("==== Welcome To Testing ===============")
    fmt.Println("=======================================")
	
}

// get user input

func main(){
	
     GreeUsers(); // function for excute 
	// declear total 
	// var conferenceTickts int = 50;
	// var remainingTickets uint = 50 ; // uint present none-negaviate values 
	// var bookings = make([]UserData,0)

	type UserData struct {  // struct is a way to group varible togetther that is have a single type
		firstName string
		lastName string
		email    string
		numberOfTickets uint
	}
	// example 
	user := UserData{
		firstName: "imkevin",
		lastName: "tido",
		email: "thisis@gmail.com",
		numberOfTickets: 30,
	}
		// user input 
	fmt.Println("Please,Input your FirstName: ");
	fmt.Scan(&user.firstName) // & pointer 
	fmt.Println("Please,Input your LastName: ")
	fmt.Scan(&user.lastName)

	fmt.Println("=============== UserDetail Information =======")
	fmt.Printf("This is UserFristName: %s\n", user.firstName)
	fmt.Printf("This is UserLastName: %s\n", user.lastName)

	fmt.Println(integers.NumberIntegers)

	// call go_packages 
	go_packages.GetWelcomeUser();
    

	fmt.Println("==============================")
}

