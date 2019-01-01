package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name + "\nWelcome on board"
	}

func sumPlusOne(a int, b int) int{
	return a + b + 1
	}

	
func main(){
	
	fmt.Println(greeting("Dred"))
	fmt.Println(sumPlusOne(2,3))
	
	}
