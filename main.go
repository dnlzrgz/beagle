package main

import (
	"log"

	"github.com/danielkvist/beagle/cmd"
)

func main() {
	root := cmd.Root()
	if err := root.Execute(); err != nil {
		log.Fatalf("error while executing the root command: %v", err)
	}
}
