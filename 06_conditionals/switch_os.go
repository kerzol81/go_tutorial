package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	
	var os string = runtime.GOOS
	
	switch os {
	
	case "darwin":
		fmt.Println("OS X.")
	
	case "linux":
		fmt.Println("Linux.")
	
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}