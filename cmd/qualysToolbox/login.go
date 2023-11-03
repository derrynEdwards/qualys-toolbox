package qualysToolbox

import (
	"fmt"
	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

func commandLogin(cfg *config.Config, apiCfg *config.ApiConfig, args ...string) error {
	msg, err := cfg.QualysApiClient.QualysLogin(apiCfg.BaseURL, apiCfg.User, apiCfg.Password)
	if err != nil {
		fmt.Println("Error on Login:", err)
		return err
	}

	fmt.Println(msg)
	return nil
}
