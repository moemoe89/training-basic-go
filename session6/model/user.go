package model

type UserForm struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

type UserListModel struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserDetailModel struct {
	Data struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserCreateModel struct {
	Data struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserUpdateModel struct {
	Data struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserDeleteModel struct {
	Data struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}