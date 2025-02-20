package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Age       int32
	CreatedAt time.Time
}
