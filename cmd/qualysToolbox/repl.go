package qualysToolbox

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, ...string) error
}

func StartRepl(cfg *config.Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("qualys-toolbox > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, args...)

			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the toolbox",
			callback:    commandExit,
		},
	}
}