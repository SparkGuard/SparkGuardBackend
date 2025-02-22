package db

import "time"

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"username"`
	Email       string `json:"email"`
	AccessLevel int    `json:"access_level"`
}

type Student struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID *uint  `json:"user_id,omitempty"`
}

type Group struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	StudentIDs []uint `json:"students,omitempty"`
	UserIDs    []uint `json:"users,omitempty"`
}

type Event struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	GroupID     uint      `json:"group"`
}

type Work struct {
	ID        uint      `json:"id"`
	Time      time.Time `json:"time"`
	StudentID uint      `json:"student"`
	EventID   uint      `json:"event"`
}

type Adoption struct {
	ID     uint    `json:"id"`
	WorkID uint    `json:"work"`
	Path   *string `json:"path"`

	PartOffset *uint `json:"part_offset,omitempty"`
	PartSize   *uint `json:"part_size,omitempty"`

	RefersTo *uint `json:"refers_to,omitempty"`
}

type Runner struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"-"`
	Tag   string `json:"tag"`
}

type Tasks struct {
	ID     uint   `json:"id"`
	WorkID uint   `json:"work"`
	Tag    string `json:"tag"`
	Status uint   `json:"status"`
}

type Error string

func (err Error) Error() string {
	return string(err)
}

const ErrNotFound = Error("not found")
