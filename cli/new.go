package jndacli

import (
	"context"
	"log"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/kadriandev/jnda/database"
	"github.com/kadriandev/jnda/model"
	"github.com/urfave/cli/v3"
)

var (
  title string
  desc string
  due int
)

var addForm =  huh.NewForm(
  huh.NewGroup(
    huh.NewInput().Title("Task Name").Inline(true).Value(&title),
    huh.NewInput().Title("Task Description").Inline(true).Value(&desc),
  ),
  huh.NewGroup(
    huh.NewSelect[int]().
      Title("Due").
      Options(
        huh.NewOption("None", 0),
        huh.NewOption("End of Day", 1),
        huh.NewOption("Tomorrow", 2),
        huh.NewOption("End of Month", 3),
      ).
      Value(&due),
  ),
);

func add() *cli.Command {
  return &cli.Command {
    Name:  "new",
    Usage: "Add a new task.",
    Action: func(ctx context.Context, cmd *cli.Command) error {

      err := addForm.Run()
      if err != nil {
          log.Fatal(err)
      }
      
      task := model.Task{
        Title:       title,
        Description: desc,
        Status:      "pending",
        DueDate:     time.Time{},
        UpdatedAt:   time.Now(),
        CreatedAt:   time.Now(),
      }
      id, err := database.AddTask(task)
      if err != nil {
        log.Panic(err)
      }
      log.Printf("Task %d added.", id)

      return nil
    },
  }
}



