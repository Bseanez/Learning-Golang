package main

import (
	"fmt"
)

func main() {
	name := "Brandon"

	fmt.Println("Hello, my name is Brandon.")   //this is how print inside of go and adds \n at the end of the line
	fmt.Printf("Hello, my name is %v.\n", name) // no new line after you print, allows you to dynamically format your output
	fmt.Print("Hello, my name is Brandon.\n")   //does not give you an extra line after printing

}
