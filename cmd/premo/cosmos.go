package main

import "github.com/urfave/cli/v2"

var cosmosCMD = &cli.Command{
	Name:  "cosmos",
	Usage: "test cosmos transaction",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "concurrent",
			Aliases: []string{"c"},
			Value:   100,
			Usage:   "concurrent number",
		},
		&cli.IntFlag{
			Name:    "tps",
			Aliases: []string{"t"},
			Value:   500,
			Usage:   "all tx number",
		},
		&cli.IntFlag{
			Name:    "duration",
			Aliases: []string{"d"},
			Value:   60,
			Usage:   "test duration",
		},
		&cli.StringFlag{
			Name:    "key_path",
			Aliases: []string{"k"},
			Usage:   "Specify key path",
		},
		&cli.StringSliceFlag{
			Name:    "remote_cosmos_addr",
			Aliases: []string{"r"},
			Usage:   "Specify remote cosmos address",
			Value:   cli.NewStringSlice("localhost:60011"),
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "Specify tx type: interchain, transfer",
			Value: "transfer",
		},
	},
	Action: cosmosBenchmark,
}

func cosmosBenchmark(ctx *cli.Context) error {
	panic("to be implemented")
}