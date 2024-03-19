package entity

import "fmt"

type User struct {
	ID       int    `json:"user_id,omitempty" db:"id"`
	UserName string `json:"username,omitempty" db:"username"`
	Balance  int    `json:"balance,omitempty" db:"balance"`
}

type UserInput struct {
	UserName string `json:"username,omitempty"`
}

func (u *UserInput) Validate() error {
	if len(u.UserName) < 2 {
		return fmt.Errorf("Слишком короткое имя пользователя")
	}
	if len(u.UserName) > 20 {
		return fmt.Errorf("Слишком длинное имя пользователя")
	}
	return nil
}
