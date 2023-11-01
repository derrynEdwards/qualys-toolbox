package main

import (
	"fmt"
	"time"

	"github.com/derrynEdwards/qualys-toolbox/cmd/qualysToolbox"
	"github.com/derrynEdwards/qualys-toolbox/internal/qualysapi"
	"github.com/derrynEdwards/qualys-toolbox/pkg/config"
)

func main() {
	qualysclient := qualysapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config.Config{
		QualysApiClient: qualysclient,
	}

	fmt.Println("===================")
	fmt.Println("  Qualys ToolBox!")
	fmt.Println("===================")
	qualysToolbox.StartRepl(cfg)
}
