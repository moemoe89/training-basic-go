package model

import "fmt"

type UserModel struct {
	ID      int    `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	Address string `json:"address" gorm:"address"`
}

// contoh replace struct pointing nama tabel
func(u UserModel) TableName() string {
	return "user"
}

// contoh membuat fungsi validasi dari struct
func (u UserModel) Validation() error {

	if len(u.Name) < 1 {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	return nil
}

// contoh membuat fungsi
func (u UserModel) Bebas() error {

	if len(u.Name) < 1 {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	return nil
}
