package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func startREPL(c *Config) error {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input:=scanner.Text()
		output:=cleanInput(input)
		cmd, ok := commands[output[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	clean_str := strings.ToLower(text)
	substrings := strings.Fields(clean_str)
	return substrings
}


