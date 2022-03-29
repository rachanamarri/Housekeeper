package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	m "app_backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
		fmt.Println(c)
		c.BindJSON(&seeker)

		login.Email = seeker.Email
		hashPassword, err := HashPassword(seeker.Password)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			login.Password = hashPassword
			fmt.Println("Hash_Password ", hashPassword)
		}

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

		login.Email = sandp.ProviderEmail
		hashPassword, err := HashPassword(sandp.ProviderPassword)
		fmt.Println("Here I am", hashPassword)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			login.Password = hashPassword
		}

		db.Create(&login)

		db.Create(&sandp)

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
			fmt.Println("Passwords are", auth.Password, storedAuth.Password)
			err := bcrypt.CompareHashAndPassword([]byte(storedAuth.Password), []byte(auth.Password))
			if err != nil {
				fmt.Println(err)
				fmt.Println("No match")
				c.JSON(401, gin.H{"message": "Login Failed!"})

			} else {
				c.JSON(http.StatusOK, storedAuth)
				fmt.Println("successfully logged in !")

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
