package main

import (
	"fmt"
	"net/http"
	"strings"

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

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	db.AutoMigrate(&Login{})

	//creating variable using gin Web Framework to handle routing and serving HTTP requests
	//r :=gin.Default() does not work, it gives a huge runtime error
	r := gin.New()

	//Routers
	r.GET("/", home)
	r.POST("/seeker_registration", create_seeker)
	r.POST("/provider_registration", create_provider)
	r.POST("/seeker_login", login_auth)
	r.POST("/provider_login", login_auth)
	r.GET("/seeker_home", nil)
	r.GET("/provider_home", nil)
	r.GET("/service/:id", nil)

	r.Run(":8080")
}

func create_seeker(c *gin.Context) {

	var seeker Seeker
	var login Login

	c.BindJSON(&seeker)

	login.Email = seeker.Email
	login.Password = seeker.Password

	db.Create(&seeker)
	db.Create(&login)

	c.JSON(http.StatusOK, gin.H{"message": "working !!"})
	fmt.Println("successfully added an entry into seeker DB !")

}

func create_provider(c *gin.Context) {

	var provider Provider
	var login Login

	c.BindJSON(&provider)

	login.Email = provider.Email
	login.Password = provider.Password

	db.Create(&provider)
	db.Create(&login)

	c.JSON(200, provider)
	fmt.Println("successfully added an entry into provider DB !")

}

func login_auth(c *gin.Context) {

	var auth Login
	var storedAuth Login

	c.BindJSON(&auth)

	err := db.Where("Email = ?", auth.Email).First(&storedAuth).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		match := strings.Compare(auth.Password, storedAuth.Password)
		if match == 0 {
			fmt.Println("match")
			c.JSON(200, gin.H{"message": "login successful!"})
		} else {
			fmt.Println("No match")
			c.JSON(401, gin.H{"message": "Login Failed!"})
		}

	}

}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Home Page"})
}
