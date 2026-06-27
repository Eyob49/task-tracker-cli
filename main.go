package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a command (add, update, delete, list, mark-in-progress, mark-done)")
		os.Exit(1)
	}

	switch args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task")
			os.Exit(1)
		}
		description := strings.Join(args[2:], " ")
		result := Add(description)
		fmt.Println(result)
	case "update":
		fmt.Println("you called the UPDATE command")
	case "delete":
		fmt.Println("you called the DELETE command")
	case "list":
		fmt.Println("you called the LIST command")
	case "mark-in-progress":
		fmt.Println("you called the MARK-IN-PROGRESS command")
	case "mark-done":
		fmt.Println("you called the MARK-DONE command")
	default:
		fmt.Println("Unknown command")
	}
}
