package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	cnt "app_backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

//var dbName string = "test.db"
//var seeker []m.Seeker
//var snp []m.ServiceAndProvider
//var login []m.Login
//var book []m.Booking

var dB *gorm.DB

// test to view a service, given its ID
func TestMain(m *testing.M) {
	log.Println("this is TestMain func")
	returnCode := m.Run()
	log.Println("ok the test was not run")
	os.Exit(returnCode)

}

func TestHomePage(t *testing.T) {

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	c.Request, _ = http.NewRequest("GET", "/", nil)

	cnt.Home(dB)

	var response gin.H

	err := json.Unmarshal(r.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	_, exists := response["message"]

	fmt.Println("...")

	assert.Equal(t, exists, true)

}
