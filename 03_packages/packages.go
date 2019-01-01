package main

import (
	"fmt"
	"math"
)

// https://golang.org/pkg/math/

func main (){
	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Sqrt(81))
	
	//
	
	result := math.Pow(3,3)
	fmt.Println(result)
	fmt.Printf("%T\n", result)
}
