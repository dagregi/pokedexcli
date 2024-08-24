package main

import "fmt"

func helpCommand(cfg *config) error {
	fmt.Println("\nPokedex CLI")
	fmt.Println("\x1B[1;2m\nCommands:\x1B[0m")
	for _, command := range getCommands() {
		fmt.Printf(" %s:\t%s\n", command.name, command.description)
	}

	return nil
}
