package main

import (
	"strings"
)

func cleanInput(text string) []string {
	clean_str := strings.ToLower(text)
	substrings := strings.Fields(clean_str)
	return substrings
}
