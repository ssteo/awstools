package cf

import (
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:      "cloudformation",
		ShortName: "cf",
		Usage:     "print CloudFormation stacks information",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "list stacks and their resources",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name",
						Usage: "stack name",
					},
					cli.StringFlag{
						Name:  "name-substring",
						Usage: "case insensitive stack name `SUBSTRING`",
					},
					cli.StringSliceFlag{
						Name:  "tag",
						Usage: "list only stacks with the given `KEY:VALUE`, can be specified multiple times",
					},
					cli.StringSliceFlag{
						Name:  "resource-type, r",
						Usage: "print only resources of the specified `TYPE`. Can be specified multiple times. Example AWS::Lambda::Function. Implies --print-resources",
					},
					cli.BoolFlag{
						Name:  "print-tags",
						Usage: "print stack tags",
					},
					cli.BoolFlag{
						Name:  "print-parameters",
						Usage: "print stack parameters",
					},
					cli.BoolFlag{
						Name:  "print-outputs",
						Usage: "print stack outputs",
					},
					cli.BoolFlag{
						Name:  "print-resources",
						Usage: "print stack resources",
					},
				},
				Action: listStacks,
			},
			{
				Name:      "events",
				Usage:     "print events of a stack",
				ArgsUsage: "STACK",
				Action:    printStackEventsCommand,
			},
			{
				Name:      "delete",
				Usage:     "delete stack",
				ArgsUsage: "STACK",
				Action:    deleteStack,
			},
		},
	}

}
