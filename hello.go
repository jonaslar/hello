package main

import "fmt"

func main() {

	a := 0
	fmt.Printf("Hello, world\n")
	fmt.Printf("Hei, alle sammen!\n")
	fmt.Printf("Dette er en test!\n")
	fmt.Printf("Dette er en test2\n")
	fmt.Print("Dette er en test3\n")

	for a < 10 {
		fmt.Print("Hei Mina!\n")
		a = a + 1
	}
}
