package qualysToolbox

import (
	"fmt"
	"os"
	"time"

	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

func commandExit(cfg *config.Config, apiCfg *config.ApiConfig, args ...string) error {
	fmt.Println("Exiting now...")
	err := cfg.QualysApiClient.QualysLogout(apiCfg.BaseURL)
	if err != nil {
		fmt.Println("Failed logging out:", err)
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Logged out successfully!")
	os.Exit(0)
	return nil
}
