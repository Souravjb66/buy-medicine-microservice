package routes
import(
	"log"
	"html/template"
	"net/http"
)
func BuyItems(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/buyitems.html")
	if err!=nil{
		log.Println(err)
	}
	temp.Execute(w,nil)
}
func BuyerIndex(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/buyerindex.html")
	if err!=nil{
		log.Println(err)
	}
	temp.Execute(w,nil)
}
