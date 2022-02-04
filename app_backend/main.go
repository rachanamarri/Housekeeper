package main

import (
	"fmt"
	"net/http"
	"strconv"
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

// type Provider struct {
// 	ServiceName      string `json:"ServiceName"`
// 	ServiceId        uint   `json:"ServiceId"`
// 	ProviderEmail    string `json:"ProviderEmail"`
// 	ProviderPassword string `json:"ProviderPassword"`
// }

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type Service struct {
// 	ServiceId          uint   `json:"ServiceId"`
// 	ProviderEmail      string `json:"ProviderEmail"`
// 	ServiceName        string `json:"ServiceName"`
// 	ServicePrice       int64  `json:"ServicePrice"`
// 	ServiceDescription string `json:"ServiceDescription"`
// }

type ServiceAndProvider struct {
	ServiceId          int64  `json:"ServiceId"`
	ServiceName        string `json:"ServiceName"`
	ProviderEmail      string `json:"ProviderEmail"`
	ProviderPassword   string `json:"ProviderPassword"`
	ServicePrice       int64  `json:"ServicePrice"`
	ServiceDescription string `json:"ServiceDescription"`
}

type Booking struct {
	ServiceId   int64  `json:"ServiceId"`
	SeekerName  string `json:"SeekerName"`
	SeekerEmail string `json:"SeekerEmail"`
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
	//db.AutoMigrate(&Provider{})
	db.AutoMigrate(&Login{})
	//db.AutoMigrate(&Service{})
	db.AutoMigrate(&Booking{})
	db.AutoMigrate(&ServiceAndProvider{})

	//creating variable using gin Web Framework to handle routing and serving HTTP requests
	//r :=gin.Default() does not work, it gives a huge runtime error
	r := gin.New()

	//Routers
	r.GET("/", home)
	r.POST("/seeker_registration", create_seeker)   //passed commandline test
	r.POST("/service_registration", create_service) //ids have to be auto-generated
	r.POST("/seeker_login", login_auth)             //passed curl test
	r.POST("/provider_login", login_auth)           //passes curl test
	r.GET("/seeker_home", nil)
	r.GET("/provider_home", nil)
	r.GET("/services", listing_services) //passed : curl http://localhost:8080/services
	r.GET("/services/:ServiceId", list_service)
	//When the seeker tries to book a service, the data has to be updated in the bookings table
	r.POST("/services/:ServiceId/book", book) //no such column error

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

func create_service(c *gin.Context) {

	//var provider Provider
	var login Login
	//var service Service
	var sandp, sandp1 ServiceAndProvider

	c.BindJSON(&sandp)

	count := db.Find(&sandp1)

	sandp.ServiceId = count.RowsAffected + 1

	// provider.ServiceName = sandp.ServiceName
	// provider.ServiceId = sandp.ServiceId
	// provider.ProviderEmail = sandp.ProviderEmail
	// provider.ProviderPassword = sandp.ProviderPassword

	login.Email = sandp.ProviderEmail
	login.Password = sandp.ProviderPassword

	// service.ServiceName = sandp.ServiceName
	// service.ServicePrice = sandp.ServicePrice
	// service.ServiceDescription = sandp.ServiceDescription
	// service.ProviderEmail = sandp.ProviderEmail

	//db.Create(&provider)
	db.Create(&login)
	//db.Create(&service)
	db.Create(&sandp)
	fmt.Println(sandp.ServiceName)

	c.JSON(200, sandp)
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

func listing_services(c *gin.Context) {

	var services []ServiceAndProvider
	if err := db.Find(&services).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, services)
	}

}

func list_service(c *gin.Context) {

	id := c.Params.ByName("ServiceId")
	fmt.Println(id)
	var service ServiceAndProvider

	if err := db.First(&service, "service_id = ?", id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, service)
	}
}

func book(c *gin.Context) {

	var booking Booking

	id, _ := strconv.ParseInt(c.Params.ByName("ServiceId"), 10, 64)
	booking.ServiceId = id
	booking.SeekerEmail = c.Params.ByName("SeekerEmail")
	booking.SeekerName = c.Params.ByName("SeekerName")

	db.Create(&booking)

	c.JSON(200, booking)

}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Home Page"})
}
