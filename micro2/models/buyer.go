package models


type Buyer struct{
	Id int `json:"id" gorm:"primary_Key"`
	First_name string `json:"first_name" gorm:"not null"`
	Last_name string `json:"last_name" gorm:"not null"`
	Email_id string `json:"email_id" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}