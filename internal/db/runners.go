package db

import (
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

// GetRunners retrieves all runners from the database.
func GetRunners() (runners []*Runner, err error) {
	runners = make([]*Runner, 0)

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "token", "tag").From("runners")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var runner Runner
		if err = rows.Scan(&runner.ID, &runner.Name, &runner.Token, &runner.Tag); err != nil {
			return
		}
		runners = append(runners, &runner)
	}

	return runners, err
}

// GetRunner retrieves a single runner by its ID.
func GetRunner(id uint) (runner *Runner, err error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "token", "tag").From("runners").Where(sb.Equal("id", id))

	query, args := sb.Build()
	row := db.QueryRow(query, args...)

	runner = &Runner{}
	if err = row.Scan(&runner.ID, &runner.Name, &runner.Token, &runner.Tag); err != nil {
		return nil, err
	}

	return runner, nil
}

// CreateRunner inserts a new runner into the database.
func CreateRunner(runner *Runner) error {
	runner.Token = uuid.New().String()

	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()
	sb.InsertInto("runners").
		Cols("name", "token", "tag").
		Values(runner.Name, runner.Token, runner.Tag).
		SQL("RETURNING id")

	query, args := sb.Build()
	return db.QueryRow(query, args...).Scan(&runner.ID)
}

// EditRunner updates a runner's details in the database.
func EditRunner(runner *Runner) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update("runners").
		Set(
			sb.Assign("name", runner.Name),
			sb.Assign("tag", runner.Tag),
		).
		Where(sb.Equal("id", runner.ID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}

// DeleteRunner deletes a runner from the database by its ID.
func DeleteRunner(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	sb.DeleteFrom("runners").Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}
