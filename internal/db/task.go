package db

import (
	"github.com/huandu/go-sqlbuilder"
	"log"
)

// TODO: Make all as transactions

func GetTasks() ([]*Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "work_id", "tag", "status").From("tasks")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	tasks := make([]*Task, 0)

	for rows.Next() {
		task := Task{}

		if err = rows.Scan(&task.ID, &task.WorkID, &task.Tag, &task.Status); err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func GetTask(id uint) (*Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "work_id", "tag", "status").From("tasks").Where(sb.Equal("id", id))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	task := Task{}

	if err = rows.Scan(&task.ID, &task.WorkID, &task.Tag, &task.Status); err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *Task) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("tasks").
		Cols("work_id", "tag", "status").
		Values(task.WorkID, task.Tag, task.Status).
		SQL("RETURNING id")

	query, args := sb.Build()
	return db.QueryRow(query, args...).Scan(&task.ID)
}

func GetTaskFromQueueForRunner(tag string) (*Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "work_id", "tag", "status").
		From("tasks").
		Where(
			sb.And(
				sb.Equal("tag", tag),
				sb.Equal("status", "In queue"),
			),
		).
		SQL("FOR UPDATE SKIP LOCKED") // Ensures only one process can access the task

	query, args := sb.Build()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
			}
		}
	}()

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	task := Task{}
	if err = rows.Scan(&task.ID, &task.WorkID, &task.Tag, &task.Status); err != nil {
		return nil, err
	}

	// Lock the task by updating the status to "Processing"
	sbUpdate := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sbUpdate.Update("tasks").
		Set(sbUpdate.Assign("status", "In work")).
		Where(sbUpdate.Equal("id", task.ID))

	updateQuery, updateArgs := sbUpdate.Build()
	_, err = tx.Exec(updateQuery, updateArgs...)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func CloseTask(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("tasks").
		Set(sb.Assign("status", "Completed")).
		Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
