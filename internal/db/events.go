package db

import (
	"github.com/huandu/go-sqlbuilder"
)

func GetEvents() (events []*Event, err error) {
	events = make([]*Event, 0)

	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.Select("id", "name", "description", "date", "group_id").From("events")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return
	}

	for rows.Next() {
		var event Event

		if err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.GroupID); err != nil {
			return
		}

		events = append(events, &event)
	}

	return events, err
}

func GetEvent(id uint) (event *Event, err error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "description", "date", "group_id").From("events").Where(sb.Equal("id", id))

	query, args := sb.Build()
	row := db.QueryRow(query, args...)

	event = &Event{}
	if err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.GroupID); err != nil {
		return nil, err
	}

	return event, nil
}

func CreateEvent(event *Event) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("events").
		Cols("name", "description", "date", "group_id").
		Values(event.Name, event.Description, event.Date, event.GroupID).SQL("RETURNING id")

	query, args := sb.Build()
	return db.QueryRow(query, args...).Scan(&event.ID)
}

func EditEvent(event *Event) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("events").
		Set(
			sb.Assign("name", event.Name),
			sb.Assign("description", event.Description),
			sb.Assign("date", event.Date),
			sb.Assign("group_id", event.GroupID),
		).
		Where(sb.Equal("id", event.ID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}

func DeleteEvent(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()

	sb.DeleteFrom("events").Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}
