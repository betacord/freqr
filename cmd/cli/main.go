// An entry point of the freqr.

package main

import (
	"log"

	"krzywicki.pro/freqr/internal/cli"
	"krzywicki.pro/freqr/internal/config"
)

func main() {
	if err := config.Initialize("config/config.yaml"); err != nil {
		log.Fatalf("Error while initializing config: %v", err)
	}

	if err := cli.Execute(); err != nil {
		log.Fatalf("Error while starting CLI app: %v", err)
	}
}
