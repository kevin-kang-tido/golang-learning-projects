package main

import (
	"fmt"
	"go_project/go_packages"
	"go_project/integers"
	testingpackage "go_project/testing_package"
	"github.com/gin-gonic/gin"
)

func main(){
     
    // this is me 


    // Create a default gin router
    r := gin.Default()

    // Define a route for the GET request at the root path "/"
    r.GET("/", func(c *gin.Context) {

        c.JSON(200, gin.H{
            "message": "Hello World",
            "product": "iphone19-pro-max",
            "gty":"200000",
            "category":"phone",
            "price":200,
    
        })
    })

    // Run the server on the default port 8080
    r.Run()














     







	fmt.Println("==============================")

     fmt.Println("this is values from testing package: ",testingpackage.Three)
    //  call from integers package
	fmt.Println("this is valuse form integers  package: ",integers.NumberIntegers)

	// call go_packages 
	go_packages.GetWelcomeUser();
    
	fmt.Println("==============================")
}

