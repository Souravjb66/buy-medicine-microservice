package routes
import(
	"log"
	"net/http"
	"html/template"
)
func RenderSignup(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/usersignup.html")
	if err!=nil{
		log.Println("render html1 problem")
	}
	temp.Execute(w,nil)
}

func RenderSaveProduct(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/saveproduct.html")
	if err!=nil{
		log.Println("render html2 problem")

	}
	temp.Execute(w,nil)
}
func RenderIndex(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/firstindex.html")
	if err!=nil{
		log.Println("render html3 problem")

	}
	temp.Execute(w,nil)

}