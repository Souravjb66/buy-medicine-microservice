package main

import (
	"log"
	"micro/micro2/database"
	"micro/micro2/routes"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
func main(){
	database.BuyerConnect()
	db:=database.Mdb.Db
	defer func(){
		ndb,err:=db.DB()
		if err!=nil{
			log.Println(err)
		}
		ndb.Close()
	}()
    wg:=sync.WaitGroup{}
	router:=mux.NewRouter()
	SaveMethod(router,&wg)
	GetHtmlPage(router)
    
	headers:=handlers.AllowedHeaders([]string{"Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","HEAD","POST","PUT","OPTIONS"})
    origins:=handlers.AllowedOrigins([]string{"*"})
    wg.Wait()

	log.Fatal(http.ListenAndServe(":8082",handlers.CORS(headers,methods,origins)(router)))

	
}
func SaveMethod(router *mux.Router,wg *sync.WaitGroup){
	wg.Add(1)
	go func(){
		defer wg.Done()
		route:=router.PathPrefix("/buyer").Subrouter()
	    route.HandleFunc("/save-buyer",routes.SaveSignup).Methods("POST")
	    route.HandleFunc("/buy-product",routes.BuyProduct).Methods("POST")
	    route.HandleFunc("/get-product",routes.GetBuyerProduct).Methods("POST")


	}()
	

}
func GetHtmlPage(router *mux.Router){
	router.HandleFunc("/buyitems",routes.BuyItems)
	router.HandleFunc("/show-buyer-items",routes.BuyerIndex)

}