package gateway
import(
	"fmt"
	"net/http"
	"io"

)

//html pages
func HtmlPageIndex(w http.ResponseWriter,r *http.Request){
	const u2="http://localhost:8081/index"
	res,err:=http.Get(u2)
	if err!=nil{
		fmt.Println(err)
	}
	defer res.Body.Close()
	w.Header().Set("Content-Type","text/html;charset-utf-8")
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		fmt.Println(rr)
	}
	
}
func HtmlPageSaveProduct(w http.ResponseWriter,r *http.Request){
	const u3="http://localhost:8081/product-store"
	res,err:=http.Get(u3)
	if err!=nil{
		fmt.Println(err)
	}
	defer res.Body.Close()
	w.Header().Set("Content-Type","text/html;charset-utf-8")
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		fmt.Println(rr)
	}

}
func HtmlPageUserSignUp(w http.ResponseWriter,r *http.Request){
	const u4="http://localhost:8081/signup"
	res,err:=http.Get(u4)
	if err!=nil{
		fmt.Println(err)
	}
	defer res.Body.Close()
	w.Header().Set("Content-Type","text/html;charset-utf-8")
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		fmt.Println(rr)
	}
}
func HtmlPageBuyItems(w http.ResponseWriter,r *http.Request){
	const u5="http://localhost:8082/buyitems"
	res,err:=http.Get(u5)
	if err!=nil{
		fmt.Println(err)

	}
	defer res.Body.Close()
	w.Header().Set("Content-Type","text/html;charset-utf-8")
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		fmt.Println(rr)
	}
}
func HtmlPageShowBuyerItems(w http.ResponseWriter,r *http.Request){
	const u6="http://localhost:8082/show-buyer-items"
	res,err:=http.Get(u6)
	if err!=nil{
		fmt.Println(err)
	}
	res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		fmt.Println(rr)
	}
}

