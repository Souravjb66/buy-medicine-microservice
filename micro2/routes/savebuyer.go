package routes
import(
	"log"
	db "micro/micro2/database"
	"net/http"
	mdware "micro/micro2/middleware"
	"micro/micro2/models"
	"encoding/json"
	"strconv"
)

func SaveSignup(w http.ResponseWriter,r *http.Request){
	base:=db.Mdb.Db
	buyer:=models.Buyer{
		First_name: r.FormValue("first_name"),
		Last_name: r.FormValue("last_name"),
		Email_id: r.FormValue("email_id"),
		Password: r.FormValue("password"),
	}
	prob:=mdware.CheckBuyerExist(&buyer.Email_id,base)
	if !prob{
		log.Println("buyer exist")
		json.NewEncoder(w).Encode("buyer exist")
		return

	}
	hashpassword:=mdware.ConvertToHash(&buyer.Password)
	if hashpassword==""{
		log.Println("hash not created")
		json.NewEncoder(w).Encode("hash not created")
		return
	}
	buyer.Password=hashpassword
	login:=models.Login{
		Email_id: buyer.Email_id,
		Password: buyer.Password,
	}
	base.Create(&buyer)
	base.Create(&login)
	json.NewEncoder(w).Encode("buyer is created")


}
func BuyProduct(w http.ResponseWriter,r *http.Request){
	pdb:=db.Mdb.Db
	email:=r.FormValue("email_id")
	password:=r.FormValue("password")
	price,calc:=strconv.Atoi(r.FormValue("price"))
	if calc!=nil{
		log.Println(calc)
	}
	login,pass:=mdware.CheckBuyerLogin(&email,&password,pdb)
	if login.Email_id=="" && !pass{
		log.Println("wrong email and password")
		json.NewEncoder(w).Encode("wrong password and email")
		return

	}
	if !pass{
		log.Println("wrong password")
		json.NewEncoder(w).Encode("wrong password")
		return

	}
	medicine:=models.Items{
		Product_name: r.FormValue("product_name"),
		Price:uint(price),
	}
	allbuy:=[]models.Items{}
	allbuy=append(allbuy, medicine)
	login.Item=allbuy
	pdb.Save(&login)
	json.NewEncoder(w).Encode("you bought it")


}
