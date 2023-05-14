package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"artist"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRoles struct {
	ID     string `json:"id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}
