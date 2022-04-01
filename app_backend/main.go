package main

import (
	s "app_backend/controllers"
	m "app_backend/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {

	//creating Database using gorm(an ORM which simplifies the mapping and persistance of the models to the database)
	db, err = gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&m.Seeker{})
	db.AutoMigrate(&m.Login{})
	db.AutoMigrate(&m.Booking{})
	db.AutoMigrate(&m.ServiceAndProvider{})

	//creating variable using gin Web Framework to handle routing and serving HTTP requests
	//r :=gin.Default() does not work, it gives a huge runtime error
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"accept", "origin", "X-Requested-With", "Content-Type", "Authorization", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", s.Home(db))
	r.POST("/seeker_registration", s.Create_seeker(db))
	r.POST("/service_registration", s.Create_service(db))
	r.POST("/seeker_login", s.Login_auth((db)))
	r.POST("/provider_login", s.Login_auth(db))
	r.GET("/seeker_home", nil)
	r.GET("/provider_home", nil)
	r.GET("/services", s.Listing_services(db))
	r.GET("/services/:ServiceId", s.List_service(db))
	//When the seeker tries to book a service, the data has to be updated in the bookings table
	r.POST("/services/:ServiceId/book", s.Book(db))

	r.Run(":8080")

}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
