package main

import (
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrateLogger struct {
	*log.Logger
	verbose bool
}

func (m migrateLogger) Verbose() bool {
	return m.verbose
}

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
					Name:     "config, c",
					Usage:    "Load configuration from `FILE`",
					Required: true,
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

				server := judge.New()

				dec := json.NewDecoder(f)

				err = dec.Decode(server)
				if err != nil {
					return err
				}

				return server.Run()
			},
		},
		{
			Name:  "web",
			Usage: "start a web server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "config, c",
					Usage:    "Load configuration from `FILE`",
					Required: true,
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

				server.Run()
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "run migrations",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "config,c",
					Usage:    "Load the web's configuration from `FILE`",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "up",
					Usage: "runs up migrations",
				},
				cli.BoolFlag{
					Name:  "down",
					Usage: "runs down migrations",
				},
				cli.IntFlag{
					Name:  "steps",
					Usage: "runs `x` up/down migrations depending on the positivity",
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

				server.ConnectToDB()
				driver, err := postgres.WithInstance(server.GetDB().DB, &postgres.Config{})
				m, err := migrate.NewWithDatabaseInstance("file://web/migrations", "postgres", driver)
				if err != nil {
					return err
				}

				m.Log = &migrateLogger{log.New(os.Stdout, "[migrate]", 0), false}

				if c.Bool("up") {
					err = m.Up()
					if err != nil {
						return err
					}
				} else if c.Bool("down") {
					err = m.Down()
					if err != nil {
						return err
					}
				} else if c.Int("steps") != 0 {
					err = m.Steps(c.Int("steps"))
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
		{
			Name:  "submit",
			Usage: "submit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "config, c",
					Usage:    "Load configuration from `FILE`",
					Required: true,
				},
				cli.IntFlag{
					Name:     "user, u",
					Usage:    "ID of user on behalf we make the submission",
					Required: true,
				},
				cli.StringFlag{
					Name:     "problemset, ps",
					Usage:    "Problemset of problem",
					Required: true,
				},
				cli.StringFlag{
					Name:     "problem, p",
					Usage:    "Problem",
					Required: true,
				},
				cli.StringFlag{
					Name:     "language, l",
					Usage:    "Language",
					Required: true,
				},
				cli.StringFlag{
					Name:     "file, f",
					Usage:    "File to submit",
					Required: true,
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

				src, err := ioutil.ReadFile(c.String("file"))
				if err != nil {
					return err
				}

				server.ConnectToDB()
				server.AddProblem(c.String("problem"))
				id, err := server.Submit(c.Int("user"), c.String("problemset"), c.String("problem"), c.String("language"), src)
				if err != nil {
					return err
				}

				log.Print("submission received with id ", id)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
