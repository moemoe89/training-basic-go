package model

import (
	"fmt"
	//"github.com/asaskevich/govalidator"
)

type UserModel struct {
	ID      int    `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	Address string `json:"address" gorm:"address"`

	Devices []*DeviceModel `json:"devices,omitempty" gorm:"ForeignKey:UserId"`
}

func(u UserModel) TableName() string {
	return "user"
}

// contoh implementasi validation
func(u UserModel) Validation() error {

	/*
	pengecekan data menggunakan lib govalidator
	if !govalidator.IsURL(u.Name) {
		return fmt.Errorf("nama harus url")
	}
	*/

	if len(u.Name) < 1 {
		return fmt.Errorf("nama harus diisi")
	}

	if len(u.Address) < 1 {
		return fmt.Errorf("alamat harus diisi")
	}

	return nil
}
