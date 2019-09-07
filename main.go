package main

import (
	"github.com/danielkvist/beagle/cmd"
)

func main() {
	root := cmd.Root()
	root.Execute()
}
