package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kadriandev/lazytask/database"
	"github.com/kadriandev/lazytask/model"
	"github.com/kadriandev/lazytask/notifier"
	"github.com/kadriandev/lazytask/visualizer"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "lazytask",
		Usage: "A TUI for taskwarrior",
    Commands: []*cli.Command{
      {
          Name:  "init",
          Usage: "init lazytask",
          Action: func(ctx context.Context, cmd *cli.Command) error {
              db, err := database.Initialize()
              if err != nil {
                  return fmt.Errorf("failed to initialize database: %w", err)
              }
              defer db.Close()
              
              fmt.Println("Database initialized successfully!")
              return nil
          },
      },
      {
          Name:  "all",
          Usage: "See all tasks.",
          Action: func(ctx context.Context, cmd *cli.Command) error {
            tasks, err := database.GetTasksWithStatus("pending") 
            if err != nil {
              log.Panic(err)
            }
            visualizer.ViewTasks(tasks)
            return nil
          },
      },
      {
          Name:  "add",
          Usage: "add a new task",
          Action: func(ctx context.Context, cmd *cli.Command) error {
            task := model.Task{
                Title: "Test",
                Description: "This is a test task",
                Status: "pending",
                DueDate: time.Time{},
                UpdatedAt: time.Now(),
                CreatedAt: time.Now(),
            }
            id, err := database.AddTask(task)
            if err != nil {
              log.Panic(err)
            }
            log.Printf("Task %d added.", id)
            return nil
          },
      },
      {
          Name:  "remove",
          Usage: "remove a task",
          Action: func(ctx context.Context, cmd *cli.Command) error {
              notifier.Alert("Successfully removed task.", "The task was removed.")
              return nil
          },
      },
    },
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
