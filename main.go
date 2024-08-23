package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	replStart(os.Stdin, os.Stdout)
}

func replStart(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "Pokedex > ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		fmt.Println(line)
	}
}
