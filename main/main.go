package main

import (
	//	"log"
	"fmt"
	"functions" // importing functions
)

func main() {
	// Why the main is not seeing functions.sumn ?
	// 1. The module is not visible - functions module itsef is not visible ?
	// 2. The function sumn is not visible ? -> Let me make it uppercase
	fmt.Println(functions.Sumn(10)) 
	fmt.Println(functions.Factorial(10))
	fmt.Println(functions.CheckPalindrome("radam"))
	fmt.Println(functions.CheckPalindrome("madam"))	
	fmt.Println(functions.CompareStr("arpit", "arpit"))
	fmt.Println(functions.CompareStr("arpit", "akhil"))
}
