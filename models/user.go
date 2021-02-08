package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        string "type:string;primary_key;default:uuid_generate_v4()"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User ..
type User struct {
	Base
	Name        string
	Email       string
	PhoneNumber string
	Balance     float64
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuidd, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuidd.String())
}

// UserRepository ..
type UserRepository interface {
	FindByID(id string) *User
	Save(user *User) (*User, error)
	Debit(id string, amount float64) error
	Credit(id string, amount float64) error
	FindAllUsers() ([]User, error)
	Delete(id string) error
}
