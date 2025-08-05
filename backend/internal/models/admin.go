package models

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	BaseModel
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (a *Admin) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}

func (a *Admin) CheckPassword(plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plaintext))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
