package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gobisect",
		Usage: "Generic tool for bisection to efficiently find \"the point\" where \"changes\" happened",
		UsageText: "Example: gobisect -c ./test.sh -p \"in11 in12\" -p \"in21 22\" -p \"in31 in32\" -p \"in41 in42\"\n" +
			"This will respectively call:\n" +
			"./test.sh in11 in12\n" +
			"./test.sh in21 in22\n" +
			"./test.sh in31 in32\n" +
			"./test.sh in41 in42\n" +
			"NOTE: Points are expected to be in order for bsearch to be able to bisect and find changing point,\n" +
			"if order of points causes command execution to return mixed outputs, bisection will give a result, but not necessarily a correct one.\n" +
			"Example of good input is when command results are such as: 0 0 0 0 0 1 1 1 1 1 1 1\n" +
			"Example of bad input is when command results are such as: 0 0 1 1 0 0 0 1 1 1 0 0\n",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "cmd",
				Aliases:  []string{"c"},
				Usage:    "command to run for bisection. The command should exit with status code other than 0 if the execution contains changes.",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "point",
				Aliases:  []string{"p"},
				Usage:    "points that are passed down to the command as inputs.",
				Required: true,
			},
			&cli.IntFlag{
				Name:        "workers",
				Aliases:     []string{"w"},
				Value:       1,
				Usage:       "if more than 1 worker is provided, multiple point executions will happen at once to possibly find the result faster.",
				DefaultText: "1",
			},
			&cli.BoolFlag{
				Name:        "more-weight-right",
				Aliases:     []string{"mwr"},
				Value:       false,
				Usage:       "if more weight is put on right, the binary search will prioritize going to the right first.",
				DefaultText: "false",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			bisect, err := NewBisect(&BisectParams{
				Cmd:             c.String("cmd"),
				Points:          c.StringSlice("point"),
				Workers:         int(c.Int("workers")),
				MoreWeightRight: c.Bool("more-weight-right"),
			})
			if err != nil {
				return err
			}
			return bisect.Run()
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("failed to bisect: %s\n", err.Error())
		os.Exit(1)
	}
}
