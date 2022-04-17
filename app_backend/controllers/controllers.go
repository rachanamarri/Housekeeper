package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	m "app_backend/model"
	"app_backend/services"

	"app_backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Create_seeker(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var seeker m.Seeker
		var login m.Login
		c.BindJSON(&seeker)
		fmt.Println(seeker)

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

func Create_provider(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var login m.Login
		var sandp, sandp1 m.Provider

		c.BindJSON(&sandp)

		count := db.Find(&sandp1)

		sandp.ProviderId = count.RowsAffected + 1

		login.Email = sandp.Email
		hashPassword, err := HashPassword(sandp.Password)
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

func Create_service(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var serv, serv1 m.Service

		c.BindJSON(&serv)

		count := db.Find(&serv1)

		serv.ServiceId = count.RowsAffected + 1
		db.Create(&serv)

		c.JSON(http.StatusOK, serv)
		fmt.Println("successfully added an entry into service DB !")
	}

	return gin.HandlerFunc(fn)

}

func Login_auth(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var auth m.Login
		var storedAuth m.Login
		c.BindJSON(&auth)
		// fmt.Println("Got from client", auth)
		err := db.Where("Email = ?", auth.Email).First(&storedAuth).Error
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {

			// fmt.Println(storedAuth.Password, auth.Password)
			err := bcrypt.CompareHashAndPassword([]byte(storedAuth.Password), []byte(auth.Password))
			if err != nil {
				fmt.Println(err)
				fmt.Println("No match")
				c.JSON(401, gin.H{"message": "Login Failed!"})

			} else {
				jwtToken, err2 := services.GenerateToken(auth.Email)
				if err2 != nil {
					c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
					c.Abort()
					return
				}

				c.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})
				fmt.Println("successfully logged in !")

			}

		}
	}

	return gin.HandlerFunc(fn)

}

func Listing_providers(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var providers []m.Provider
		if err := db.Find(&providers).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, providers)
		}
	}
	return gin.HandlerFunc(fn)

}

func List_service(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		id := c.Params.ByName("ProviderId")
		fmt.Println(id)
		var service m.Provider

		if err := db.First(&service, "ProviderId = ?", id).Error; err != nil {
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
		requiredToken := c.Request.Header["Authorization"]

		// Check if the token is provided
		print(len(requiredToken))
		if len(requiredToken) == 0 {
			// Abort with error
			fmt.Println("1.No session found for current user")
			c.AbortWithStatus(404)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No session found for current user"})
			return
		}
		fmt.Println("Reached here 1")
		_, err := middleware.Authenticate(requiredToken[0])
		fmt.Println("Reached here 2")
		if err != nil {
			fmt.Println("2.No session found for current user")
			c.AbortWithStatus(403)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No session found for current user"})
			return
		}

		var booking m.Booking

		c.BindJSON(&booking)

		db.Create(&booking)

		c.JSON(200, booking)
	}

	return gin.HandlerFunc(fn)

}

func ProviderDetails(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var snp m.Provider
		var rtngs m.Ratings

		id := c.Params.ByName("ServiceId")
		fmt.Println(id, snp, rtngs)

	}
	return gin.HandlerFunc(fn)
}

func Rate(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var rate m.Ratings
		var service m.Provider

		id, _ := strconv.ParseInt(c.Params.ByName("ServiceId"), 10, 64)
		rate.ServiceID = id

		if err := db.First(&service, "ServiceId = ?", id).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println("the provider does not exist")
		} else {
			rate.ProviderEmail = service.Email
		}

		rate.Rating, _ = strconv.ParseInt(c.Params.ByName("rating"), 10, 64)

		db.Create(&rate)

		c.JSON(200, rate)

	}

	return gin.HandlerFunc(fn)
}

func Home(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Home Page"})
	}

	return gin.HandlerFunc(fn)
}
