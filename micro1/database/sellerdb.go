package database
import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"micro/micro1/models"
)
type SellerDBInstance struct{
	Sldb *gorm.DB
	
}
var Sdb SellerDBInstance
func Connect(){
	dsn:="root:sourav@90###@tcp(localhost:3306)/medseller?parseTime=true"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Println("micro1 seelerdb",err)
	}
	Sdb.Sldb=db
	err=db.AutoMigrate(&models.Seller{},&models.Login{},&models.Product{})
	if err!=nil{
		log.Println("in micro1 migrate",err)
	}
}