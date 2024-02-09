package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Game struct definition
type Game struct {
	Name   string
	Setup  string
	Date   string
	Token  string
	Amount float64
}

func parseText(text string) (*Game, error) {
	game := &Game{}

	// Split the text into lines
	lines := strings.Split(text, "\n")

	// Flag to indicate whether we've encountered the first line
	var challengeStarted bool

	// Regular expression to match lines starting with "*"
	re := regexp.MustCompile(`^\s*\*(\w+):\s*(.*)$`)

	// Loop through each line
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check if the line starts with "Open /stadium Challenge:"
		if !challengeStarted {
			if line == "Open /stadium Challenge:" {
				challengeStarted = true
			}
			continue
		}

		// Check if the line starts with "*"
		if strings.HasPrefix(line, "*") {
			// Match the line with the regular expression
			match := re.FindStringSubmatch(line)
			if len(match) == 3 {
				key := strings.Title(match[1]) // Convert first letter to uppercase
				value := match[2]

				// Map the extracted field names to the corresponding struct fields
				switch key {
				case "Name":
					game.Name = value
				case "Setup":
					game.Setup = value
				case "Date":
					game.Date = value
				case "Amount":
					// Convert the amount to uint64
					amount, err := strconv.ParseUint(value, 10, 64)
					if err != nil {
						return nil, err
					}
					game.Amount = amount
				}
			}
		}
	}

	return game, nil
}

func main() {
	text := `
        Open /stadium Challenge:
               *Game: Tekken8
               *Match Set Up: Best of 5
               *Match Date: Feb 4th at 6pm
               *Wager Amount: $ 250 usdc @versus
    `

	game, err := parseText(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", game.Name)
	fmt.Println("Setup:", game.Setup)
	fmt.Println("Date:", game.Date)
	fmt.Println("Amount:", game.Amount)
}
