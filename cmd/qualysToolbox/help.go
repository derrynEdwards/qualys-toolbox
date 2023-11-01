package qualysToolbox

import (
	"fmt"

	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

func commandHelp(cfg *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Qualys Toolbox!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
