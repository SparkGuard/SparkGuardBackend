package db

import (
	"github.com/huandu/go-sqlbuilder"
	"log"
)

func GetWorks() (works []*Work, err error) {
	works = make([]*Work, 0)
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "time", "event_id", "student_id").From("works")

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

	for rows.Next() {
		var work Work
		if err = rows.Scan(&work.ID, &work.Time, &work.EventID, &work.StudentID); err != nil {
			return nil, err
		}
		works = append(works, &work)
	}

	return works, nil
}

// TODO: Add context
func GetWorksIdOfEvent(eventID uint64) (works []uint64, err error) {
	works = make([]uint64, 0)
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.
		Select("id").
		From("works").
		Where(sb.Equal("event_id", eventID))

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

	for rows.Next() {
		var id uint64
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		works = append(works, id)
	}

	return works, nil
}

func GetWork(id uint) (work *Work, err error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "time", "event_id", "student_id").From("works").Where(sb.Equal("id", id))

	query, args := sb.Build()

	row := db.QueryRow(query, args...)

	var result Work
	if err = row.Scan(&result.ID, &result.Time, &result.EventID, &result.StudentID); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateWork(work *Work) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("works").
		Cols("time", "event_id", "student_id").
		Values(work.Time, work.EventID, work.StudentID).
		SQL("RETURNING id")

	query, args := sb.Build()

	err := db.QueryRow(query, args...).Scan(&work.ID)

	go func() {
		unique_tags := sqlbuilder.PostgreSQL.NewSelectBuilder()
		unique_tags.Distinct().Select("tag").From("runners")

		sb = sqlbuilder.PostgreSQL.NewInsertBuilder()
		sb.InsertInto("tasks")
		sqlbuilder.With(sqlbuilder.CTETable("unique_tags").As(unique_tags))
	}()

	return err
}

func EditWork(work *Work) error {

	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("works").
		Set(
			sb.Assign("time", work.Time),
			sb.Assign("event_id", work.EventID),
			sb.Assign("student_id", work.StudentID),
		).
		Where(sb.Equal("id", work.ID))

	query, args := sb.Build()

	_, err := db.Exec(query, args...)

	return err
}

func DeleteWork(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()

	sb.DeleteFrom("works").
		Where(sb.Equal("id", id))

	query, args := sb.Build()

	_, err := db.Exec(query, args...)

	return err
}
