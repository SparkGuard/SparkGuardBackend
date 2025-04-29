package db

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
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

	selectQuery, selectArgs := sb.Build()

	// Begin a new transaction
	conn, err := pgx.Connect(context.TODO(), connectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.TODO())

	tx, err := conn.Begin(context.TODO())
	if err != nil {
		return nil, err
	}
	// Ensure we rollback the transaction if an error occurs
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(context.TODO()); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(context.TODO()); commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
			}
		}
	}()

	// Execute SELECT query to fetch the task
	row := tx.QueryRow(context.TODO(), selectQuery, selectArgs...)

	task := Task{}
	if err = row.Scan(&task.ID, &task.WorkID, &task.Tag, &task.Status); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	sbUpdate := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sbUpdate.Update("tasks").
		Set(sbUpdate.Assign("status", "In work")).
		Where(sbUpdate.Equal("id", task.ID))
	updateQuery, updateArgs := sbUpdate.Build()

	// Lock the task by updating its status to "In work"
	_, err = tx.Exec(context.TODO(), updateQuery, updateArgs...)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func GetAllTasksFromQueueForRunner(tag string) ([]*Task, uint, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("tasks.id", "tasks.work_id", "tasks.tag", "tasks.status", "works.event_id").
		From("tasks").
		JoinWithOption(sqlbuilder.InnerJoin, "works", "tasks.work_id = works.id").
		Where(
			sb.And(
				sb.Equal("tasks.tag", tag),
				sb.Equal("tasks.status", "In queue"),
			),
		).
		Limit(1).
		SQL("FOR UPDATE SKIP LOCKED") // Ensure rows are locked for the transaction

	selectQuery, selectArgs := sb.Build()

	// Begin a new transaction
	conn, err := pgx.Connect(context.TODO(), connectionString)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close(context.TODO())

	tx, err := conn.Begin(context.TODO())
	if err != nil {
		return nil, 0, err
	}

	// Ensure we rollback the transaction if an error occurs
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(context.TODO()); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(context.TODO()); commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
			}
		}
	}()

	row := tx.QueryRow(context.TODO(), selectQuery, selectArgs...)
	var eventID uint
	firstTask := Task{}
	if err = row.Scan(&firstTask.ID, &firstTask.WorkID, &firstTask.Tag, &firstTask.Status, &eventID); err != nil {
		if err == pgx.ErrNoRows {
			return nil, 0, ErrNotFound
		}
		return nil, 0, err
	}

	// Fetch all tasks of the same event
	sbAll := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sbAll.Select("tasks.id", "tasks.work_id", "tasks.tag", "tasks.status").
		From("tasks").
		JoinWithOption(sqlbuilder.InnerJoin, "works", "tasks.work_id = works.id").
		Where(
			sbAll.And(
				sbAll.Equal("works.event_id", eventID),
				sbAll.Equal("tasks.status", "In queue"),
			),
		).
		SQL("FOR UPDATE SKIP LOCKED")

	selectAllQuery, selectAllArgs := sbAll.Build()
	rows, err := tx.Query(context.TODO(), selectAllQuery, selectAllArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		rows.Close()
	}()

	tasks := make([]*Task, 0)
	ids := make([]interface{}, 0)
	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.ID, &task.WorkID, &task.Tag, &task.Status); err != nil {
			return nil, 0, err
		}

		tasks = append(tasks, &task)
		ids = append(ids, task.ID)
	}

	if len(tasks) == 0 {
		return nil, 0, ErrNotFound
	}

	// Lock all tasks by updating their status to "In work"
	sbUpdate := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sbUpdate.Update("tasks").
		Set(sbUpdate.Assign("status", "In work")).
		Where(sbUpdate.In("id", ids...))
	updateQuery, updateArgs := sbUpdate.Build()

	_, err = tx.Exec(context.TODO(), updateQuery, updateArgs...)
	if err != nil {
		return nil, 0, err
	}

	return tasks, eventID, nil
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

func CloseTaskWithError(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("tasks").
		Set(sb.Assign("status", "Error")).
		Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

// ResetTask resets a task's status to "In queue"
func ResetTask(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("tasks").
		Set(sb.Assign("status", "In queue")).
		Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}

// GetTasksByWorkID returns all tasks associated with a specific work_id
func GetTasksByWorkID(workID uint) ([]*Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "work_id", "tag", "status").
		From("tasks").
		Where(sb.Equal("work_id", workID))

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
