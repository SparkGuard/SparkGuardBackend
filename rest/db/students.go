package db

import "github.com/huandu/go-sqlbuilder"

func GetStudents() ([]Student, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "email", "user_id").From("students")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	students := make([]Student, 0)

	for rows.Next() {
		var student Student

		if err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.UserID); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func GetStudent(id uint) (*Student, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "email", "user_id").From("students").Where(sb.Equal("id", id))

	query, args := sb.Build()
	row := db.QueryRow(query, args...)

	var student Student

	if err := row.Scan(&student.ID, &student.Name, &student.Email, &student.UserID); err != nil {
		return nil, err
	}

	return &student, nil
}

func CreateStudent(student *Student) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("students").Cols("name", "email", "user_id").Values(student.Name, student.Email, student.UserID).SQL("RETURNING id")

	query, args := sb.Build()

	return db.QueryRow(query, args...).Scan(&student.ID)
}

func EditStudent(student *Student) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("students").Set(sb.Assign("name", student.Name), sb.Assign("email", student.Email)).Where(sb.Equal("id", student.UserID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}
