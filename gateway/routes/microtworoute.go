package gateway

import (
	"log"
	"net/http"
	"io"
	"net/url"
)
func PostSaveBuyer(w http.ResponseWriter,r *http.Request){
	const b1="http://localhost:8082/buyer/save-buyer"
	result:=url.Values{}
	result.Set("first_name",r.FormValue("first_name"))
	result.Set("last_name",r.FormValue("last_name"))
	result.Set("password",r.FormValue("password"))
	result.Set("email_id",r.FormValue("email_id"))
	res,err:=http.PostForm(b1,result)
	if err!=nil{
		log.Println(err)
	}
	defer res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}
}
func PostBuyerBuyProduct(w http.ResponseWriter,r *http.Request){
	const b2="http://localhost:8082/buyer/buy-product"
	result:=url.Values{}
	result.Set("email_id",r.FormValue("email_id"))
	result.Set("password",r.FormValue("password"))
	result.Set("price",r.FormValue("price"))
	result.Set("product_name",r.FormValue("product_name"))
	res,err:=http.PostForm(b2,result)
	if err!=nil{
		log.Println(err)
	}
	defer res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}


}
func GetBuyerProductList(w http.ResponseWriter,r *http.Request){
	const b3="http://localhost:8082/buyer/buy-product"
	result:=url.Values{}
	result.Set("email_id",r.FormValue("email_id"))
	result.Set("password",r.FormValue("password"))
	res,err:=http.PostForm(b3,result)
	if err!=nil{
		log.Println(err)
	}
	defer res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}
}