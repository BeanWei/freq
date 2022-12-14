package cmd

import (
	"github.com/BeanWei/freq/internal/app"
	"github.com/BeanWei/freq/internal/task"
	"github.com/urfave/cli/v2"
)

var (
	Start = &cli.Command{
		Name:  "start",
		Usage: "start app",
		Action: func(ctx *cli.Context) error {
			go task.Run()
			app.NewHTTPServer()
			return nil
		},
	}
)
