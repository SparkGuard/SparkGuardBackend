package db

import (
	"github.com/huandu/go-sqlbuilder"
)

func GetUsers() ([]*User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "email", "access_level").From("users")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		user := User{}

		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccessLevel); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func GetUser(id uint) (*User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "email", "access_level").From("users").Where(sb.Equal("id", id))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	user := User{}

	if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccessLevel); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *User, salt string, hash string) (err error) {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("users").Cols("name", "email", "access_level", "salt", "password").Values(user.Name, user.Email, user.AccessLevel, salt, hash).SQL("RETURNING id")

	query, args := sb.Build()
	return db.QueryRow(query, args...).Scan(&user.ID)
}

func EditUser(user *User) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("users").Set(
		sb.Assign("name", user.Name),
		sb.Assign("email", user.Email),
		sb.Assign("access_level", user.AccessLevel),
	).Where(sb.Equal("id", user.ID))

	query, args := sb.Build()
	result, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
