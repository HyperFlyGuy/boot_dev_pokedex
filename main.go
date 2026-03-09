package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
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
		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

