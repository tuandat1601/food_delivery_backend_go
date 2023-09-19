package main

import (
	"food_delivery/internal/model"
	"food_delivery/internal/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type TableName struct {
	TableName string
}

func main() {

	dsn := "host=localhost user=tuandat password=tuandat1601 dbname=mydatabase port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to POSTGRES:", err)
	}
	log.Println("Connected:", db)
	r := gin.Default()

	var tableNames []TableName
	result := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tableNames)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Xử lý danh sách tên các bảng
	var tableList string
	for _, tableName := range tableNames {
		tableList += tableName.TableName + "\n"
	}
	
	routes.UserRoutes(r,db)
	r.GET("/tables", func(c *gin.Context) {
		c.String(200, "Danh sách các bảng:\n%s", tableList)
	})
	r.POST("/auth", postAuth(db))
	r.Run()
}
func postAuth(db *gorm.DB) func (*gin.Context){
	
	return func (c *gin.Context){

		var newuser  model.User
		if err := c.ShouldBind(&newuser); err != nil {
			return
		    }

		    if err := db.Table("users").Create(&newuser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newuser})
	}

}