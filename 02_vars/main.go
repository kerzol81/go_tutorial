package main

import "fmt"

	/*
	string
	bool
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte - alias for uint8
	rune - alias for int32
	float32 float64
	complex64 complex128
	*/
	
	// variables must be used

func main(){
	var i int = 8
	var s string = "aaa"
	//const
	const isCool = true
	
	// assignment
	name, email := "ZOL", "zol@gmail.com"
	
	//shorthand only inside a function:
	n := "zol" 
	
	fmt.Println(i, s, isCool, n, name, email)
	// get the type of
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", name)
	
	}
