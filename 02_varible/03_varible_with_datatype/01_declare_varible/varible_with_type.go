package main

import "fmt"

func main(){
	
    // Declaring a string variable
    var varVariable string = "this is var variable with datatype!"
    
    // Declaring an integer variable
    var intVariable int = 42

    // Declaring a float variable
    var floatVariable float64 = 3.14159

    // Declaring a boolean variable
    var boolVariable bool = true

    // Declaring an array variable
    var arrayVariable [3]int = [3]int{1, 2, 3}

    // Declaring a slice variable
    var sliceVariable []string = []string{"apple", "banana", "cherry"}

    // Declaring a map variable
    var mapVariable map[string]int = map[string]int{"age": 30, "score": 100}

    // Declaring a pointer variable
    var pointerVariable *int = &intVariable

    // Declaring a struct variable
    type Person struct {
        Name string
        Age  int
    }
    var structVariable Person = Person{Name: "John", Age: 30}

    // Printing all variables
    fmt.Println("String variable:", varVariable)
    fmt.Println("Integer variable:", intVariable)
    fmt.Println("Float variable:", floatVariable)
    fmt.Println("Boolean variable:", boolVariable)
    fmt.Println("Array variable:", arrayVariable)
    fmt.Println("Slice variable:", sliceVariable)
    fmt.Println("Map variable:", mapVariable)
    fmt.Println("Pointer variable:", pointerVariable)
    fmt.Println("Struct variable:", structVariable)
}