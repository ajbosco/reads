package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/ajbosco/reads/goodreads"
	"github.com/ajbosco/reads/version"
)

var (
	configFile string
	debug      bool

	client *goodreads.Client
)

func main() {

	app := cli.NewApp()
	app.Name = "reads"
	app.Usage = "Command line tool to interact with Goodreads"
	app.Version = version.VERSION

	// Setup the global flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Goodreads CLI config file",
			EnvVar:      "GOODREADS_CLI_CONFIG",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "enable debug logging",
			Destination: &debug,
		},
	}

	// Build the list of available commands.
	app.Commands = []cli.Command{
		searchCommand,
		shelvesCommand,
	}

	// Set the before function.
	app.Before = func(c *cli.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
			fmt.Println("Debug logging is enabled")
		}

		if len(configFile) <= 0 {
			return errors.New("Goodreads config file cannot be empty")
		}

		cfg := new(goodreads.Config)
		cfg, err := goodreads.ReadConfig(configFile)
		if err != nil {
			logrus.Debug(err)
			return errors.New("Could not read Goodreads config file")
		}
		if cfg.DeveloperKey == "" || cfg.DeveloperSecret == "" {
			return errors.New("Goodreads config file requires Developer Key and Developer Secret")
		}

		// Get OAuth tokens if they are not in config.
		if cfg.AccessToken == "" || cfg.AccessSecret == "" {
			token, err := goodreads.GetAccessToken(cfg.DeveloperKey, cfg.DeveloperSecret)
			if err != nil {
				logrus.Debug(err)
				return errors.New("Could not get Goodreads access token")
			}
			cfg.AccessToken = token.Token
			cfg.AccessSecret = token.Secret
			goodreads.WriteConfig(cfg, configFile)
		}

		// Create the Goodreads client.
		client, err = goodreads.NewClient(cfg)
		if err != nil {
			logrus.Debug(err)
			return errors.New("Could not create Goodreads client")
		}

		return nil
	}

	app.Run(os.Args)
}
