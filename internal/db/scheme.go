package db

import "time"

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"username"`
	Email       string `json:"email"`
	AccessLevel string `json:"access_level"`
	Password    string `json:"-"`
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

type AdoptionVerdict string

const (
	AdoptionNotIssued       AdoptionVerdict = "Not Issued"
	AdoptionInsignificantly                 = "Insignificantly"
	AdoptionSignificantly                   = "Significantly"
	AdoptionBlatant                         = "Blatant"
)

type Adoption struct {
	ID     uint64  `json:"id"`
	WorkID uint64  `json:"work"`
	Path   *string `json:"path"`

	PartOffset *uint64 `json:"part_offset,omitempty"`
	PartSize   *uint64 `json:"part_size,omitempty"`
	RefersTo   *uint64 `json:"refers_to,omitempty"`

	SimilarityScore float32 `json:"similarity_score"`
	IsAIGenerated   bool    `json:"is_ai_generated"`

	Verdict     AdoptionVerdict `json:"verdict"`
	Description string          `json:"description"`
}

type Runner struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"-"`
	Tag   string `json:"tag"`
}

type Task struct {
	ID     uint   `json:"id"`
	WorkID uint   `json:"work"`
	Tag    string `json:"tag"`
	Status string `json:"status"`
}

type Error string

func (err Error) Error() string {
	return string(err)
}

const ErrNotFound = Error("not found")
