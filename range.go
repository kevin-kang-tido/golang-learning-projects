package main
import "fmt"


func Testing2() {

	fmt.Println("====================================================")

    nums :=[]int{1,2,3,4,5}

	for index, value := range nums {

		fmt.Printf("Index: %d, Value: %d\n", index, value)

	}
	fmt.Println("====================================================")
    

}