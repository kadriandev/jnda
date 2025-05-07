package jndacli

import (
	"context"
	"log"

	"github.com/kadriandev/jnda/database"
	"github.com/kadriandev/jnda/visualizer"
	"github.com/urfave/cli/v3"
)

func all() *cli.Command {
  return &cli.Command{
    Name:  "list",
    Usage: "List agenda items",
    Action: func(ctx context.Context, cmd *cli.Command) error {
      tasks, err := database.GetTasksWithStatus("pending")
      if err != nil {
        log.Panic(err)
      }
      visualizer.ViewTasks(tasks)
      return nil
    },
  }
}
