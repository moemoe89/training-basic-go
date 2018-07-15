package model

import "fmt"

type UserModel struct {
	ID      int    `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	Address string `json:"address" gorm:"address"`
}

func(u UserModel) TableName() string {
	return "user"
}

func (u UserModel) Validation() error {

	if len(u.Name) < 1 {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	return nil
}

func (u UserModel) Bebas() error {

	if len(u.Name) < 1 {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	return nil
}