package main

import "fmt"

func main(){
	// arrays has to be fixed lenght
	
	var fruits [3]string
	
	fruits[0] = "orange"
	fruits[1] = "banana"
	fruits[2] = "avocado"
	
	//fruits := [2]string{"orange", "banana"}
	
	fmt.Println(fruits)
	fmt.Println(fruits[0])
	
	fruits2 := []string{"kiwi", "banana", "apple"}
	
	fmt.Println(fruits2)
	fmt.Println(len(fruits2))
	
	//slicing:
	fmt.Println(fruits2[0:1])
	}
