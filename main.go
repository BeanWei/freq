package main

import (
	"os"

	"github.com/BeanWei/freq/internal/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Freq",
		Commands: []*cli.Command{
			cmd.Start,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
