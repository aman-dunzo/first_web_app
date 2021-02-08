package models

import (
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindByID(id string) *User {
	var user User
	r.db.First(&user, "id = ?", id)

	return &user
}

func (r *UserRepo) Save(user *User) (*User, error) {
	result := r.db.Create(&user)

	return user, result.Error
}

func (r *UserRepo) Debit(id string, amount float64) error {
	var user User
	r.db.First(&user, "id = ?", id)
	user.Balance -= amount
	r.db.Save(user)

	return r.db.Error
}

func (r *UserRepo) Credit(id string, amount float64) error {
	var user User
	r.db.First(&user, "id = ?", id)
	user.Balance += amount
	r.db.Save(user)

	return r.db.Error
}

func (r *UserRepo) FindAllUsers() ([]User, error) {
	users := make([]User, 0)
	r.db.Find(&users)

	return users, r.db.Error
}

func (r *UserRepo) Delete(id string) error {
	r.db.Delete(&User{}, "id = ?", id)

	return r.db.Error
}
