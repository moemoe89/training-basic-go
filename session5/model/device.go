package model

type DeviceModel struct {
	ID     int       `json:"id" gorm:"id"`
	UserID int       `json:"user_id" gorm:"user_id"`
	Device string    `json:"device" gorm:"device"`

	User   *UserModel `json:"user,omitempty" gorm:"User"`
}

func(u DeviceModel) TableName() string {
	return "device"
}
