package main

import(
	"fmt"
	"os"
)
type cliCommand struct{
	name string
	description string
	callback func(c *Config, args []string) error
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
			name: "mapb",
			description: "Will display the names of the previous 20 locations in the Pokemon world. If you are on the first page it should display 'you are on the first page'",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "This command take the name of a location area and lists all the pokemon currently located here.",
			callback: commandExplore,
		},
	}
}


func commandExit(c *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, args []string) error{
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _,command := range commands{
		fmt.Printf("\n%s: %s\n",command.name,command.description)
	}
	return nil
}

func commandMap(c *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		url= *c.next
	}
	data := c.client.LocationAreaRequest(url)
	c.previous = data.Previous
	c.next = data.Next
	for _,res := range data.Results{
		fmt.Println(res.Name)
	}
	return nil
}
func commandMapb(c *Config, args []string) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url:= *c.previous
	data := c.client.LocationAreaRequest(url)
	c.previous = data.Previous
	c.next = data.Next
	for _,res := range data.Results{
		fmt.Println(res.Name)
	}
	return nil
}

func commandExplore(c *Config, args []string) error {
	if len(args) == 0{
		fmt.Println("No location was given")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	data := c.client.ExploreRequest(url)
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _,res := range data.PokemonEncounters{
		fmt.Printf("  - %s\n",res.Pokemon.Name)
	}
	return nil
}
