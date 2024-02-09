package main

import (
	"errors"
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

func main() {
	text1 := `
		Open /stadium Challenge:
			*Game: Tekken8
			*Match Set Up: Best of 5
			*Match Date: Feb 4th at 6pm
			*Wager Amount: $ 250 usdc @versus
	`
	game, err := parseText(text1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	_ = game
}

func parseText(text string) (*Game, error) {
	game := &Game{}

	re := regexp.MustCompile(`^\s*\*\s*(.+?):\s*(.*)$`)

	mentionRe := regexp.MustCompile(`@\w+`)

	lines := strings.Split(text, "\n")
	if len(lines) >= 15 {
		str := fmt.Sprintf("lines of number %v exceeds Maximum Number of lines: 15 .",
			len(lines))
		return nil, errors.New(str)
	}

	var challengeStarted bool

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if !challengeStarted {
			if line == "Open /stadium Challenge:" {
				challengeStarted = true
			}
			continue
		}

		line = mentionRe.ReplaceAllString(line, "")

		if strings.HasPrefix(line, "*") {
			match := re.FindStringSubmatch(line)
			if len(match) == 3 {
				key := strings.Title(match[1])
				value := match[2]

				switch key {
				case "Game":
					game.Name = value
				case "Match Set Up":
					game.Setup = value
				case "Match Date":
					game.Date = value
				case "Wager Amount":
					wagerAmountStr := strings.TrimSpace(value)
					if !strings.HasPrefix(wagerAmountStr, "$") {
						return nil, fmt.Errorf("invalid input format: missing currency symbol in amount '%s'", wagerAmountStr)
					}

					wagerAmountParts := strings.Fields(value)
					if len(wagerAmountParts) != 3 {
						return nil, errors.New("Invalid length of amount field.")
					}

					amountStr := wagerAmountParts[1]
					token := wagerAmountParts[2]
					amount, err := strconv.ParseFloat(amountStr, 64)
					if err != nil {
						return nil, fmt.Errorf("invalid input format: invalid amount '%s'", value)
					}
					game.Amount = amount
					game.Token = token
				default:
					return nil, fmt.Errorf("invalid input format: unexpected field '%s'", key)
				}
			}
		}
	}
	if game.Name == "" || game.Setup == "" || game.Date == "" || game.Amount == 0 || game.Token == "" {
		return nil, errors.New("invalid input format: missing required fields")
	}

	return game, nil
}
