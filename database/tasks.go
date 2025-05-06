package database

import (
	"log"
	"time"
  "github.com/kadriandev/lazytask/model"
)

func GetTasksWithStatus(status string) ([]model.Task, error) {
  var tasks []model.Task 
  db, err := GetDatabaseConnection()
  if err != nil {
    return []model.Task{}, err
  }
  defer db.Close()

  rows, err := db.Query(`SELECT * FROM tasks WHERE status = ?`, status)
  if err != nil {
    log.Fatal("Error fetching tasks.")
  }
  defer rows.Close();

  for rows.Next() {
    item := model.Task{}
    err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Status, &item.DueDate, &item.CreatedAt, &item.UpdatedAt)
    if err != nil {
        log.Fatal(err)
    }
    tasks = append(tasks, item)
  }

  return tasks, nil
}

func AddTask(task model.Task) (int64, error) {
  db, err := GetDatabaseConnection()
  if err != nil {
    return 0, err
  }
  defer db.Close()

  result, err := db.Exec(`
    INSERT INTO tasks (title, description, status, due_date, created_at, updated_at) 
    VALUES (?, ?, ?, ?, ?, ?)`,
    task.Title,
    task.Description,
    task.Status,
    task.DueDate,
    task.CreatedAt,
    task.UpdatedAt,
  )
  if err != nil {
    return 0, err
  }

  // Get the ID of the inserted task
  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return id, nil
}

func UpdateTask(taskId int64, task model.Task) (int64, error) {
  db, err := GetDatabaseConnection()
  if err != nil {
    return 0, err
  }
  defer db.Close()

  result, err := db.Exec(`
    UPDATE tasks SET 
      title = ?,
      description = ?,
      status = ?,
      due_date = ?,
      updated_at = ?
    where id = ?`,
    task.Title,
    task.Description,
    task.Status,
    task.DueDate,
    time.Now(),
    taskId,
  )
  if err != nil {
    return 0, err
  }

  // Get the ID of the updated task
  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return id, nil
}

func DeleteTask(taskId int64) (int64, error) {
  db, err := GetDatabaseConnection()
  if err != nil {
    return 0, err
  }
  defer db.Close()

  result, err := db.Exec(`DELETE FROM tasks where id = ?`, taskId)
  if err != nil {
    return 0, err
  }

  // Get the ID of the deleted task
  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return id, nil
}
