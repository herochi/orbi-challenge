package domain

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Age       int32              `bson:"age"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is empty")
	}

	if u.Age < 0 {
		return errors.New("age can't be less than zero")
	}

	if u.Email == "" {
		return errors.New("email is empty")
	}

	if !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return errors.New("field email doesn't have email format")
	}

	return nil
}

type UpdateUser struct {
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (u *UpdateUser) Validate() error {
	if u.Email != " " {
		if !regexp.MustCompile(emailRegex).MatchString(u.Email) {
			return errors.New("field email doesn't have email format")
		}
	}

	return nil
}
