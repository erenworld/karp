package main

import (
	"fmt"
	"os"

	"karp/repl"
)

func main() {
	fmt.Printf("Welcome! This is the Karp programming language!\n")
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}