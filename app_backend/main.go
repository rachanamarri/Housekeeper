package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Seeker struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type Provider struct {
	ServiceName string `json:"serviceName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

var db *gorm.DB
var err error

func main() {

	//creating Database using gorm(an ORM which simplifies the mapping and persistance of the models to the database)
	db, err = gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Seeker{})
	db.AutoMigrate(&Provider{})

	//creating variable using gin Web Framework to handle routing and serving HTTP requests
	//r :=gin.Default() does not work, it gives a huge runtime error
	r := gin.New()

	//Routers
	r.GET("/", home)
	r.POST("/seeker_registration", create_seeker)
	r.POST("/provider_registration", create_provider)
	r.POST("/seeker_login", nil)
	r.POST("/provider_login", nil)
	r.GET("/seeker_home", nil)
	r.GET("/provider_home", nil)

	r.Run(":8080")
}

func create_seeker(c *gin.Context) {

	var seeker Seeker
	c.BindJSON(&seeker)
	db.Create(&seeker)
	c.JSON(200, seeker)
	fmt.Println("successfully added an entry into seeker DB !")

}

func create_provider(c *gin.Context) {

	var provider Provider
	c.BindJSON(&provider)
	db.Create(&provider)
	c.JSON(200, provider)
	fmt.Println("successfully added an entry into provider DB !")

}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Home Page"})
}
