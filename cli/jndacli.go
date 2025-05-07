package jndacli

import "github.com/urfave/cli/v3"

func App(version string) *cli.Command {
  return &cli.Command{
      Name:    "jnda",
      Version: version,
      Usage:   "Smart session manager for the terminal",
      Commands: []*cli.Command{
        all(),
        add(),
      },
	}
}
