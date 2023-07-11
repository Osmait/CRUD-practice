package main

import (
	cmd "github.com/osmait/crud/cmd/bosstrap"

	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
