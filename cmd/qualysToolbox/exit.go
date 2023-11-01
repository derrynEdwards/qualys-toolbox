package qualysToolbox

import (
	"fmt"
	"os"
	"time"

	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

func commandExit(cfg *config.Config, args ...string) error {
	fmt.Println("Exiting now...")
	time.Sleep(1 * time.Second)
	os.Exit(0)
	return nil
}
