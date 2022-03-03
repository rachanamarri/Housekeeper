package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	m "app_backend/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Create_seeker(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var seeker m.Seeker
		var login m.Login

		c.BindJSON(&seeker)

		login.Email = seeker.Email
		login.Password = seeker.Password

		db.Create(&seeker)
		db.Create(&login)

		c.JSON(http.StatusOK, gin.H{"message": "working !!"})
		fmt.Println("successfully added an entry into seeker DB !")
	}

	return gin.HandlerFunc(fn)

}

func Create_service(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var login m.Login
		var sandp, sandp1 m.ServiceAndProvider

		c.BindJSON(&sandp)

		count := db.Find(&sandp1)

		sandp.ServiceId = count.RowsAffected + 1

		login.Email = sandp.ProviderEmail
		login.Password = sandp.ProviderPassword

		db.Create(&login)

		db.Create(&sandp)
		fmt.Println(sandp.ServiceName)

		c.JSON(200, sandp)
		fmt.Println("successfully added an entry into provider DB !")
	}

	return gin.HandlerFunc(fn)

}

func Login_auth(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var auth m.Login
		var storedAuth m.Login

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

	return gin.HandlerFunc(fn)

}

func Listing_services(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var services []m.ServiceAndProvider
		if err := db.Find(&services).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, services)
		}
	}
	return gin.HandlerFunc(fn)

}

func List_service(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		id := c.Params.ByName("ServiceId")
		fmt.Println(id)
		var service m.ServiceAndProvider

		if err := db.First(&service, "service_id = ?", id).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, service)
		}
	}

	return gin.HandlerFunc(fn)
}

func Book(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var booking m.Booking

		id, _ := strconv.ParseInt(c.Params.ByName("ServiceId"), 10, 64)
		booking.ServiceId = id
		booking.SeekerEmail = c.Params.ByName("SeekerEmail")
		booking.SeekerName = c.Params.ByName("SeekerName")

		db.Create(&booking)

		c.JSON(200, booking)
	}

	return gin.HandlerFunc(fn)

}

func Home(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Home Page"})
	}

	return gin.HandlerFunc(fn)
}
