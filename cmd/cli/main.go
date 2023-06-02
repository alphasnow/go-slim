package main

import (
	"go-slim/internal/cmds"
	"log"
)

func main() {
	err := cmds.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
