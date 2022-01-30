package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
)

type Seeker struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func main() {

	//creating Database using gorm(an ORM which simplifies the mapping and persistance of the models to the database)
	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Seeker{})

	//creating variable using gin Web Framework to handle routing and serving HTTP requests
	r := gin.New()

	r.GET("/", h)
	r.Run(":8080")
}

func h(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is the main page of our HouseKeeper Website. Here only two login/register buttons will be displayed saperately for user and service provider. This selection will redirect them to the next pages/routes....."})
}
