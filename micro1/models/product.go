package models

type Product struct{
	Id int `json:"id" gorm:"primary_Key"`
	Product_name string `json:"product_name" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`
	Medicine_type string `json:"medicine_type" gorm:"not null"`
    LoginId int `json:"loginid" gorm:"not null"`
}