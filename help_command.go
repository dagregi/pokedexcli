package main

import "fmt"

func helpCommand() error {
	fmt.Println("\nPokedex CLI")
	fmt.Println("\nCommands:")
	for _, command := range getCommands() {
		fmt.Printf(" %s:\t%s\n", command.name, command.description)
	}

	return nil
}
