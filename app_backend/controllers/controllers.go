package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	m "app_backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", 12090292)
	session.Set("email", "test@gmail.com")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}

func Create_seeker(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var seeker m.Seeker
		var login m.Login

		c.BindJSON(&seeker)

		login.Email = seeker.Email
		login.Password = seeker.Password

		db.Create(&seeker)
		db.Create(&login)

		c.JSON(http.StatusOK, seeker)
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
		fmt.Println(sandp.ProviderEmail, sandp.ProviderPassword)
		login.Email = sandp.ProviderEmail
		login.Password = sandp.ProviderPassword

		db.Create(&login)

		db.Create(&sandp)
		fmt.Println(sandp.ServiceName)

		c.JSON(http.StatusOK, sandp)
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
				// fmt.Println("Reached here 6")
				// fmt.Println("match")
				// session := sessions.Default(c)
				// session.Set("id", 12090292)
				// session.Set("email", "test@gmail.com")
				// session.Save()
				c.JSON(http.StatusOK, storedAuth)
				fmt.Println("successfully logged in !")
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

		if err := db.First(&service, "ServiceId = ?", id).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, service)
			//fix it : sensitive info like password sould not be displayed on the browser
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

func ProviderDetails(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var snp m.ServiceAndProvider
		var rtngs m.Ratings

		id := c.Params.ByName("ServiceId")
		fmt.Println(id)

	}
	return gin.HandlerFunc(fn)
}

func Home(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Home Page"})
	}

	return gin.HandlerFunc(fn)
}
