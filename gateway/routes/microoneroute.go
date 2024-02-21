package gateway

import (
	"log"
	"io"
	"net/http"
	"net/url"
)

//data
func GetAllProductData(w http.ResponseWriter,r *http.Request){
	const g1="http://localhost:8081/get/allproducts"
	res,err:=http.Get(g1)
	if err!=nil{
		log.Println(err)
	}
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}
   
	defer res.Body.Close()
}
func GetSellerProduct(w http.ResponseWriter,r *http.Request){
	const g2="http://localhost:8081/get/sellerproduct"
	res,err:=http.Get(g2)
	if err!=nil{
		log.Println(err)
	}
	defer res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}

}
func PostSellerSignUp(w http.ResponseWriter,r *http.Request){
	const p1="http://localhost:8081/seller/save-user"
	formdata:=url.Values{}
	formdata.Set("first_name",r.FormValue("first_name"))
	formdata.Set("last_name",r.FormValue("last_name"))
	formdata.Set("password",r.FormValue("password"))
	formdata.Set("email_id",r.FormValue("email_id"))

	res,err:=http.PostForm(p1,formdata)
	if err!=nil{
		log.Println(err)
	
	}
	defer res.Body.Close()
    _,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}

}
func PostSellerSaveProduct(w http.ResponseWriter,r *http.Request){
	const p2="http://localhost:8081/seller/save-product"
	formdata:=url.Values{}
	formdata.Set("product_name",r.FormValue("product_name"))
	formdata.Set("medicine_type",r.FormValue("medicine_type"))
	formdata.Set("prce",r.FormValue("price"))
	formdata.Set("email_id",r.FormValue("email_id"))
	formdata.Set("password",r.FormValue("password"))
	res,err:=http.PostForm(p2,formdata)
	if err!=nil{
		log.Println(err)
	}
	defer res.Body.Close()
	_,rr:=io.Copy(w,res.Body)
	if rr!=nil{
		log.Println(rr)
	}

}