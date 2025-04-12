package user

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthdate string    `json:"birthdate"`
	CreatedAt time.Time `json:"created_at"`
}

type admin struct {
	email    string
	password string
	user     user
}

func New(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastName, birthdate is required")
	}

	return &user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

func (user *user) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *user) ClearUserDetail() {
	user.firstName = ""
	user.lastName = ""
	user.birthdate = ""
}
