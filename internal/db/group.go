package db

import "github.com/huandu/go-sqlbuilder"

func LoadGroupUsers(group *Group) error {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("user_id").From("group_users").Where(sb.Equal("group_id", group.ID))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return err
	}

	defer rows.Close()

	group.UserIDs = make([]uint, 0)

	for rows.Next() {
		var id uint

		if err = rows.Scan(&id); err != nil {
			return err
		}

		group.UserIDs = append(group.UserIDs, id)
	}

	return nil
}

func LoadGroupStudents(group *Group) error {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("user_id").From("group_users").Where(sb.Equal("group_id", group.ID))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return err
	}

	defer rows.Close()

	group.StudentIDs = make([]uint, 0)

	for rows.Next() {
		var id uint

		if err = rows.Scan(&id); err != nil {
			return err
		}

		group.StudentIDs = append(group.StudentIDs, id)
	}

	return nil
}

func GetGroups() ([]*Group, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name").From("groups")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var groups = make([]*Group, 0)

	for rows.Next() {
		var group Group

		if err = rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}

		groups = append(groups, &group)
	}

	return groups, nil
}

func GetGroup(id uint) (group *Group, err error) {
	group = new(Group)

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name").From("groups").Where(sb.Equal("id", id))

	query, args := sb.Build()

	if err = db.QueryRow(query, args...).Scan(&group.ID, &group.Name); err != nil {
		return nil, err
	}

	if err = LoadGroupUsers(group); err != nil {
		return nil, err
	}

	if err = LoadGroupStudents(group); err != nil {
		return nil, err
	}

	return
}

func CreateGroup(group *Group) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("groups").Cols("name").Values(group.Name).SQL("RETURNING id")

	query, args := sb.Build()

	return db.QueryRow(query, args...).Scan(&group.ID)
}

func EditGroup(group *Group) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("groups").Set(sb.Assign("name", group.Name)).Where(sb.Equal("id", group.ID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

func AddUserToGroup(groupID, userID uint) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("group_users").Cols("group_id", "user_id").Values(groupID, userID)

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

func RemoveUserFromGroup(groupID, userID uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()

	sb.DeleteFrom("group_users").Where(sb.Equal("group_id", groupID), sb.Equal("user_id", userID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

func AddStudentToGroup(groupID, studentID uint) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.InsertInto("group_students").Cols("group_id", "student_id").Values(groupID, studentID)

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

func RemoveStudentFromGroup(groupID, studentID uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()

	sb.DeleteFrom("group_students").Where(sb.Equal("group_id", groupID), sb.Equal("student_id", studentID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

func DeleteGroup(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()

	sb.DeleteFrom("groups").Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}
