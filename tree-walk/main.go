package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: tree-walk.exe [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		fmt.Println("Running file: " + os.Args[1] + "")
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("got from file:\n" + string(bytes))
	//run(string(bytes))
}

func runPrompt() {
	var line string
	fmt.Scanln(line)
	fmt.Println("got " + line)
}
