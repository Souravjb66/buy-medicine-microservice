package middleware
import(
	"log"
	"gorm.io/gorm"
	"micro/micro2/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckBuyerExist(email *string,db *gorm.DB)bool{
	buyer:=models.Buyer{}
	err:=db.Where("email_id=?",*email).Find(&buyer)
	if err!=nil{
		log.Println(err)
	}
	if *email!=buyer.Email_id{
		return true

	}
	return false

}
func ConvertToHash(password *string)string{
	bytepassword:=[]byte(*password)
	hashpassword,er:=bcrypt.GenerateFromPassword(bytepassword,bcrypt.DefaultCost)
	if er!=nil{
		log.Println("in hash password",er)
		return ""

	}
	return string(hashpassword)


}
func CheckBuyerLogin(email *string,password *string,db *gorm.DB)(models.Login,bool){
	login:=models.Login{}
	bytepassword:=[]byte(*password)
	ok:=db.Where("email_id=?",*email).Find(&login)
	if ok!=nil{
		log.Println(ok)
	}
	err:=bcrypt.CompareHashAndPassword([]byte(login.Password),bytepassword)
	if err!=nil{
		log.Println("passord is incorrect")
		return login,false
	}
	return login,true


}