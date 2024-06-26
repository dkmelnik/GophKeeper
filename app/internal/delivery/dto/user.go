package dto

import (
	"fmt"
	"strings"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

type ChangePassword struct {
	PreviousPassword string
	NewPassword      string
	ConfirmPassword  string
}

func (r *ChangePassword) Validate() error {
	if len(strings.TrimSpace(r.NewPassword)) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	if r.ConfirmPassword != r.ConfirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}
	return nil
}

type Login struct {
	Login    string
	Password string
}

func (r *Login) Validate() error {
	if len(strings.TrimSpace(r.Login)) < 2 {
		return fmt.Errorf("login must be at least 2 characters long")
	}

	if len(strings.TrimSpace(r.Password)) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	return nil
}

type Register struct {
	Login           string
	Password        string
	ConfirmPassword string
}

func (r *Register) Validate() error {
	if len(strings.TrimSpace(r.Login)) < 2 {
		return fmt.Errorf("login must be at least 2 characters long")
	}

	if len(strings.TrimSpace(r.Password)) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}

	if r.Password != r.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}

	return nil
}

func (r *Register) ToEntity() entity.User {
	return entity.User{
		Login:    r.Login,
		Password: r.Password,
	}
}
