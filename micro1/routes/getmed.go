package routes
import(
	"log"
	"net/http"
	"micro/micro1/database"
	"micro/micro1/models"
	"encoding/json"
	mdware "micro/micro1/middleware"

)
func GetAllProduct(w http.ResponseWriter,r *http.Request){
	db:=database.Sdb.Sldb
	products:=[]models.Product{}
	err:=db.Find(&products)
	if err!=nil{
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
func GetSellerProducts(w http.ResponseWriter,r *http.Request){
	db:=database.Sdb.Sldb
	email:=r.FormValue("email_id")
	password:=r.FormValue("password")
	data,prob:=mdware.CheckLoginCorrect(&email,&password,db)
	if data.Email_id=="" && !prob{
		json.NewEncoder(w).Encode("wrong email & password")
		return
	}else if data.Email_id!="" && !prob{
		json.NewEncoder(w).Encode("wrong password")
		return

	}
	json.NewEncoder(w).Encode(data)

}

