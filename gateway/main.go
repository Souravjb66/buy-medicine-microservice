package main

import (
	// "encoding/json"
	"fmt"
	// "io"
	"log"
	routes "micro/gateway/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router:=mux.NewRouter()
	fmt.Println("in main")
	
    HtmlRenderMicroOne(router)
	headers:=handlers.AllowedHeaders([]string{"Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","HEAD","POST","PUT","OPTIONS"})
    origins:=handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(router)))
}
func HtmlRenderMicroOne(router *mux.Router,){
	next:=router.PathPrefix("/gatewayseller").Subrouter()
	router.HandleFunc("/saveproduct",routes.HtmlPageSaveProduct).Methods("GET")
	next.HandleFunc("/index",routes.HtmlPageIndex).Methods("GET")
	next.HandleFunc("/usersignup",routes.HtmlPageUserSignUp).Methods("GET")
	next.HandleFunc("/buyitems",routes.HtmlPageBuyItems).Methods("GET")
	next.HandleFunc("/showbuyeritems",routes.HtmlPageShowBuyerItems).Methods("GET")
	
}
func GetMicroOne(router *mux.Router){
	next:=router.PathPrefix("/getseller").Subrouter()
	next.HandleFunc("/allproducts",routes.GetAllProductData).Methods("GET")
	next.HandleFunc("/sellerproducts",routes.GetSellerProduct).Methods("GET")


}
func PostMicroOne(router *mux.Router){
	next:=router.PathPrefix("/postseller").Subrouter()
	next.HandleFunc("/signup",routes.PostSellerSignUp).Methods("POST")
	next.HandleFunc("/saveproduct",routes.PostSellerSaveProduct).Methods("POST")

}
func PostMicroTwo(router *mux.Router){
	next:=router.PathPrefix("/buyer").Subrouter()
	next.HandleFunc("/signup",routes.PostSaveBuyer).Methods("POST")
	next.HandleFunc("/buyproduct",routes.PostBuyerBuyProduct).Methods("POST")
	next.HandleFunc("/productlist",routes.GetBuyerProductList).Methods("POST")
}