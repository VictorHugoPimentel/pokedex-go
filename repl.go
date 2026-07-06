package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please Enter some Text: ")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println("You entered:", text)
	}

}

func cleanInput(str string)[]string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}