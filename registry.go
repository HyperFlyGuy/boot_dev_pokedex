package main

import(
	"fmt"
	"os"
)
type cliCommand struct{
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Will display the names of 20 locations in the Pokemon world. Each subsequent call will display the next 20 locations and so on",
			callback: commandMap,
		}
	}
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _,command := range commands{
		fmt.Printf("\n%s: %s\n",command.name,command.description)
	}
	return nil
}

func commandMap() error{

}
