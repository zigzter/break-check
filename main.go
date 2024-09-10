package main

import (
	"log"

	"github.com/zigzter/break-check/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
