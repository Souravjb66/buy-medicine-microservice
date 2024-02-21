package routes
import(
	"log"
	"net/http"
	"micro/micro1/models"
	"micro/micro1/database"
	mdware "micro/micro1/middleware"
	"encoding/json"
	"strconv"
)
func SaveSeller(w http.ResponseWriter,r *http.Request){ //seller create 
	db:=database.Sdb.Sldb
	seller:=models.Seller{
		First_name: r.FormValue("first_name"),
		Last_name: r.FormValue("last_name"),
		Password: r.FormValue("password"),
		Email_id: r.FormValue("email_id"),
	}
	prob:=mdware.CheckSellerExist(&seller.Email_id,db)
	if !prob{
		log.Println("email already exist!!!")
		return
	}
	hashpassword:=mdware.CreateHashPassword(&seller.Password)
	if hashpassword==""{
		json.NewEncoder(w).Encode("hash password not created")
		return
	}
	seller.Password=hashpassword
	login:=models.Login{
		Email_id: seller.Email_id,
		Password: seller.Password,
	}
	db.Create(&seller)
	db.Create(&login)
	json.NewEncoder(w).Encode(seller.First_name)

	
}
func SaveProducts(w http.ResponseWriter,r *http.Request){
	db:=database.Sdb.Sldb
	value,err:=strconv.Atoi(r.FormValue("price"))
	if err!=nil{
		log.Println("price not converted to int!!",err)
	}
	products:=models.Product{
		Product_name: r.FormValue("product_name"),
		Price:uint(value),
		Medicine_type: r.FormValue("medicine_type"),


	}
	email:=r.FormValue("email_id")
	password:=r.FormValue("password")
	data,prob:=mdware.CheckLoginCorrect(&email,&password,db)
	if data.Email_id==""{
		json.NewEncoder(w).Encode("incorrect email ")
		return
	}else if data.Email_id!="" && !prob{
		json.NewEncoder(w).Encode("incorrect password")
		return
	}
	var allproducts []models.Product
	allproducts=append(allproducts,products)
	data.Allproducts=allproducts
	db.Save(&data)
	json.NewEncoder(w).Encode(products.Medicine_type)


}
