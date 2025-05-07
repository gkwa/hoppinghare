package main

import (
	"os"

	"github.com/gkwa/hoppinghare/cmd"
	"github.com/gkwa/hoppinghare/internal/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error("Error: %s", err)
		os.Exit(1)
	}
}

