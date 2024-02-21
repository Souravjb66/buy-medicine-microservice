package database

import(
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"micro/micro2/models"

)
type BuyerInstance struct{
	Db *gorm.DB
}
var Mdb BuyerInstance
func BuyerConnect(){
	dsn:="root:sourav@90###@tcp(localhost:3306)/micro2buyer?parseTime=true"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Println(err)
	}
	Mdb.Db=db
	problem:=db.AutoMigrate(&models.Buyer{},&models.Login{},&models.Items{})
	if problem!=nil{
		log.Println(problem)
	}
	


}
