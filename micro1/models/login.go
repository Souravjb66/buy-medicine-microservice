package models

type Login struct{
	Id int `json:"id" gorm:"primary_Key"`
	Password string `json:"password" gorm:"not null"`
	Email_id string `json:"email_id" gorm:"unique"`
	Allproducts []Product `json:"allproducts" gorm:"foreignKey:LoginId"`
}