package model

type User struct {
	Model
	Username  string `gorm:"uniqueIndex;size:30;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   string `json:"addresss"`
	// Roles     []Role
}

type Role struct {
	Model
	Title string
}
