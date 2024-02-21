package models

type Items struct{
	Id int `json:"id" gorm:"primary_Key"`
	Product_name string `json:"product_name"`
	Price uint `json:"price"`
	Client []Login `json:"client" gorm:"many2many:client_product"`
}