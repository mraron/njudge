package main

import (
	"encoding/json"
	"errors"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "njudge"
	app.Usage = "CLI utility for njudge"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "judge",
			Usage: "start a judging server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Load configuration from `FILE`",
				},
			},
			Action: func(c *cli.Context) error {
				name := c.String("config")
				if len(name) == 0 {
					return errors.New("config file is required")
				}

				f, err := os.Open(name)
				if err != nil {
					return err
				}

				server := &judge.Server{}

				dec := json.NewDecoder(f)

				err = dec.Decode(server)
				if err != nil {
					return err
				}

				return judge.NewFromCloning(server).Run()
			},
		},
		{
			Name:  "web",
			Usage: "start a web server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Load configuration from `FILE`",
				},
			},
			Action: func(c *cli.Context) error {
				name := c.String("config")
				if len(name) == 0 {
					return errors.New("config file is required")
				}

				f, err := os.Open(name)
				if err != nil {
					return err
				}

				server := &web.Server{}

				dec := json.NewDecoder(f)

				err = dec.Decode(server)
				if err != nil {
					return err
				}

				return server.Run()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
