package models

type User struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:name`
	Email    string `json:"email" grom:"unique"`
	Password []byte `json:"password"`
}

type ServicesUser struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:name`
	Email    string `json:"email" grom:"unique"`
	Password []byte `json:"password"`
}
