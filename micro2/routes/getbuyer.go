package routes
import(
	"log"
	maindb "micro/micro2/database"
	"micro/micro2/models"
	mdware "micro/micro2/middleware"
	"net/http"
	"encoding/json"
)

func GetBuyerProduct(w http.ResponseWriter,r *http.Request){
	db:=maindb.Mdb.Db
	email:=r.FormValue("email_id")
	password:=r.FormValue("password")
	exist,pass:=mdware.CheckBuyerLogin(&email,&password,db)
	if exist.Email_id =="" && !pass{
		log.Println("wrong email and passowrd")
		json.NewEncoder(w).Encode("wrong email and password")
		return
	}else if exist.Email_id!="" && !pass{
		log.Println("wrong password")
		json.NewEncoder(w).Encode("wrong password")
		return
	}
	user:=models.Login{}
	ok:=db.Where("email_id=?",email).Find(&user)
	if ok!=nil{
		log.Println(ok)
	}

	sec:=db.Model(&user).Association("Item").Find(&user.Item)
	if sec!=nil{
		log.Println(sec)
	}
	
	json.NewEncoder(w).Encode(user)
	

}