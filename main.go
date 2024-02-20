package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wandersonpaes/check-ip-host-command-line/app"
)

func main() {
	fmt.Println("Starting")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
