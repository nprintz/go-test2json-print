package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func print() {
	fmt.Print("no trailing newline")
}

func println() {
	fmt.Println("has a trailing newline")
}

