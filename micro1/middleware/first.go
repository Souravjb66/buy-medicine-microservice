package middleware

import (
	"log"
	"micro/micro1/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckSellerExist(email *string,db *gorm.DB)bool{
	type seller struct{
		Id int `json:"id" gorm:"primary_Key"`
	    First_name string `json:"first_name" gorm:"not null"`
	    Last_name string `json:"last_name" gorm:"not null"`
	    Password string `json:"password" gorm:"not null"`
	    Email_id string `json:"email_id" gorm:"unique"`
	}
	var slr seller
    err:=db.Where("email_id=?",*email).Find(&slr)
	if err!=nil{
		log.Println(err)
	}
	if *email!=slr.Email_id{
		return true
	}
	return false
	

}
func CheckLoginCorrect(email *string,password *string,db *gorm.DB)(models.Login,bool){
	slr:=models.Login{}
	err:=db.Where("email_id=?",*email).Find(&slr)
	if err!=nil{
		log.Println(err)
	}
	bytepassword:=[]byte(*password)
	er:=bcrypt.CompareHashAndPassword([]byte(slr.Password),bytepassword)
	if er!=nil{
		log.Println("email ;;;",slr.Email_id)
		
		if slr.Email_id=="" {
			log.Println(" wrong email!!")
			return slr,false

		}
		log.Println("password not match!",err)
		return slr,false
	}
	return slr,true

}
func CreateHashPassword(password *string)string{
	bytepassword:=[]byte(*password)
	result,err:=bcrypt.GenerateFromPassword(bytepassword,bcrypt.DefaultCost)
	if err!=nil{
		log.Println("hash not created!!",err)
		return ""
	}
	return string(result)

}