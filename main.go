package main

import (
	"fmt"
)

// Requirements:
// Each product name has at least one abbreviation in the acronym
// Each abbreviation in the acronym has a respective product name
// The letters in the acronym appear in order for the entire string of product names

func main() {
	acronymChars := []string{"test", "test"}
	fmt.Println(pop(acronymChars))
}

func Pop(strings []string) (string, []string) {
	popped, strings := strings[len(strings)-1], strings[:len(strings)-1]
	return popped, strings
}
