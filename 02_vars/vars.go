package main

import "fmt"

	/* MAIN TYPES
	
	string
	bool
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte - alias for uint8
	rune - alias for int32
	float32 float64
	complex64 complex128
	
	*/
	
	// variables must be  used, unless it will not compile

func main(){
	var i int = 10
	var s string = "AAAAAAA"
	//const
	const isCool = true
	
	// assignment, go will know that these are strings
	name, email := "zol", "zol@gmail.com"
	
	//shorthand; only inside a function:
	n := "zol" 
	
	//just to make sure to I used all variables:
	fmt.Println(i, s, isCool, n, name, email)
	
	// get the type of:
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", name)
	
	}
