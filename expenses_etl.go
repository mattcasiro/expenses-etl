package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	var input string
	var output string
	var verbosity int
	max_verbosity := 2

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "input",
			Aliases:     []string{"i"},
			Usage:       "Load input data from `FILEPATH`",
			Required:    true,
			Destination: &input,
			Category:    "IO",
			// Verify file exists
			Action: func(ctx *cli.Context, s string) error {
				if _, err := os.Stat(s); os.IsNotExist(err) {
					return fmt.Errorf("file does not exist: %s", s)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Save output data to `FILEPATH`",
			Required:    true,
			Destination: &output,
			Category:    "IO",
			// Verify file exists
			Action: func(ctx *cli.Context, s string) error {
				if _, err := os.Stat(s); os.IsNotExist(err) {
					return fmt.Errorf("file does not exist: %s", s)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Usage:    "Load YAML transform configuration from `FILEPATH`",
			Required: true,
			Category: "IO",
			// Verify file exists
			Action: func(ctx *cli.Context, s string) error {
				if _, err := os.Stat(s); os.IsNotExist(err) {
					return fmt.Errorf("file does not exist: %s", s)
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:     "verbose",
			Usage:    "How verbose the program output should be",
			Aliases:  []string{"v"},
			Count:    &verbosity,
			Category: "Misc",
			// Verify there are no more than one instance of this flag
			Action: func(ctx *cli.Context, b bool) error {
				if ctx.Count("verbose") > max_verbosity {
					verbosity = max_verbosity
					log.Printf("Warning: Verbosity is capped at %d, additional instances have no effect.", max_verbosity)
				}
				return nil
			},
		},
	}
	app := &cli.App{
		Name:                   "Expenses Transformer",
		Usage:                  "Transform expenses data from bank download to usable CSV.",
		UseShortOptionHandling: true,
		Flags:                  flags,
		Before:                 altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
		Action: func(cCtx *cli.Context) error {
			if verbosity > 0 {
				fmt.Println("Verbosity:", verbosity)
			}
			name := "hardcoded_name"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().First()
			}
			if input == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
