package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "count_cli"
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop,s",
					Usage: "value to stop counter",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				if stop <= 0 {
					fmt.Println("Stop can't be a negative")
				} else {
					for i := 1; i <= stop; i++ {
						fmt.Println(i)
					}

				}
				return nil

			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "count down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Start couner down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start <= 0 {
					fmt.Println("Start can't be a negative")
				} else {
					for i := start; i >= 0; i-- {
						fmt.Println(i)
					}

				}
				return nil

			},
		},
	}

	app.Run(os.Args)

}
