package models

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	Id           uint   `json:"id"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
}

func (admin *Admin) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	admin.Password = hashPassword
}

func (admin *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(admin.Password, []byte(password))

}
