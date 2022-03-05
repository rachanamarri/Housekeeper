package main

import (
	"log"
	"os"
	"testing"
)

//var dbName string = "test.db"
//var seeker []m.Seeker
//var snp []m.ServiceAndProvider
//var login []m.Login
//var book []m.Booking

// test to view a service, given its ID
func TestMain(m *testing.M) {
	log.Println("this is TestMain func")
	returnCode := m.Run()
	log.Println("ok the test was not run")
	os.Exit(returnCode)

}

func TestHomePage(t *testing.T) {
	log.Println("this is my firstt test")
}
