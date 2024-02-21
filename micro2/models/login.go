package models

type Login struct{
	Id int `json:"id" gorm:"primary_Key"`
	Email_id string `json:"email_id" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Item []Items `json:"item" gorm:"many2many:client_product"`
}