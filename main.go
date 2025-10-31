package main

import (
	"karp/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
