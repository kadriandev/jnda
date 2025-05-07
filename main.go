package main

import (
	"context"
	"log"
	"os"

	jndacli "github.com/kadriandev/jnda/cli"
)

func main() {
  version := "0.1"
	if err := jndacli.App(version).Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
