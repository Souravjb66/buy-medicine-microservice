package main

import (
	"fmt"
	"log"
	"micro/micro1/database"
	"micro/micro1/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func main() {
	database.Connect()
	sldb:=database.Sdb.Sldb
	defer func(){
		db,err:=sldb.DB()
		if err!=nil{
			log.Panic("db not closed")
		}
		db.Close()
	}()
    wg:=sync.WaitGroup{}

	router:=mux.NewRouter()
    fmt.Println("micro1")
	Saveis(router,&wg)
	Getis(router,&wg)
	GetHtml(router)
	headers:=handlers.AllowedHeaders([]string{"Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","HEAD","POST","PUT","OPTIONS"})
    origins:=handlers.AllowedOrigins([]string{"*"})
	wg.Wait()

	http.ListenAndServe(":8081",handlers.CORS(headers,methods,origins)(router))
}
func Saveis(router *mux.Router,wg *sync.WaitGroup){
	wg.Add(1)
	go func(){
		defer wg.Done()
		secend:=router.PathPrefix("/seller").Subrouter()
	    secend.HandleFunc("/save-user",routes.SaveSeller).Methods("POST")
	    secend.HandleFunc("/save-product",routes.SaveProducts).Methods("POST")

	}()
	

}
func Getis(router *mux.Router,wg *sync.WaitGroup){
	wg.Add(1)
	go func(){
		defer wg.Done()
		second:=router.PathPrefix("/get").Subrouter()
	    second.HandleFunc("/allproducts",routes.GetAllProduct).Methods("GET")
	    second.HandleFunc("/sellerproduct",routes.GetSellerProducts).Methods("POST")

	}()
	

}
func GetHtml(router *mux.Router){

	router.HandleFunc("/signup",routes.RenderSignup).Methods("GET")
	router.HandleFunc("/product-store",routes.RenderSaveProduct).Methods("GET")
	router.HandleFunc("/index",routes.RenderIndex).Methods("GET")
}