package main

import (
	"fmt"
	"food_delivery/docs"
	"food_delivery/internal/routes"
	"log"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type TableName struct {
	TableName string
}

// @BasePath 

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
     }
func main() {
	err:= godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	port :=os.Getenv("PORT")
	postgresUser :=os.Getenv("POSTGRES_USER")
	postgresPort :=os.Getenv("POSTGRES_PORT")
	postgresHost :=os.Getenv("POSTGRES_HOST")
	postgresPassword :=os.Getenv("POSTGRES_PASSWORD")
	postgresDB :=os.Getenv("POSTGRES_DB")
	dsn :=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",postgresHost,postgresUser,postgresPassword,postgresDB,postgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to POSTGRES:", err)
	}
	log.Println("Connected:", db,port)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
 
	var tableNames []TableName
	result := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tableNames)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	var tableList string
	for _, tableName := range tableNames {
		tableList += tableName.TableName + "\n"
	}
	
	routes.UserRoutes(r,db)
	routes.RestaurantRoutes(r,db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/tables", func(c *gin.Context) {
		c.String(200, "Danh sách các bảng:\n%s", tableList)
	})
	
	r.Run((":"+port))
}
