package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	app := app.CreateApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
