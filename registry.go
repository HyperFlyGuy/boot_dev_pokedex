package main

import(
	"fmt"
	"os"
)
type cliCommand struct{
	name string
	description string
	callback func(c *Config) error
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
			description: "Will display the names of the next 20 locations in the Pokemon world. Each subsequent call will display the next 20 locations and so on",
			callback: commandMap,
		},
		"mapb": {
			name: "map",
			description: "Will display the names of the previous 20 locations in the Pokemon world. If you are on the first page it should display 'you are on the first page'",
			callback: commandMapb,
		},
	}
}


func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error{
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _,command := range commands{
		fmt.Printf("\n%s: %s\n",command.name,command.description)
	}
	return nil
}

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		url= *c.next
	}
	data := PokeRequest(url)
	c.previous = data.Previous
	c.next = data.Next
	for _,res := range data.Results{
		fmt.Println(res.Name)
	}
	return nil
}
func commandMapb(c *Config) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url:= *c.previous
	data := PokeRequest(url)
	c.previous = data.Previous
	c.next = data.Next
	for _,res := range data.Results{
		fmt.Println(res.Name)
	}
	return nil
}
