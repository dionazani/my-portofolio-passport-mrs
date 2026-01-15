package infrastructure_entities

import "time"

type AppPerson struct {
	ID          string
	Fullname    string
	Email       string
	MobilePhone string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
