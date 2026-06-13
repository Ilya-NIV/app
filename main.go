package main

import "fmt"

func main() {
	fmt.Println("Enter your target, and I will predict an active molecule: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Macrocycle is the best solution for %s!", name)
}
