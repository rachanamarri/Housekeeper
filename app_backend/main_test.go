package main

import (
	s "app_backend/controllers"
	m "app_backend/model"
	"fmt"
	"strconv"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	// "gotest.tools/assert"
)

type APIEnv struct {
	DB *gorm.DB
}

type ConsolePrinter struct{}

//utility function
func (cp *ConsolePrinter) Print(value string) {
	fmt.Printf("this is value: %s", value)
}

func TestCreateSeekerAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_seeker := m.Seeker{
		Name:     "Nancy",
		Email:    "nancy123@gmail.com",
		Password: "123",
		Address:  "7512 SW 34th STREET GNV FL 32608",
	}

	reqBody, err := json.Marshal(mock_seeker)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateSeekerRouter(db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Seeker{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := mock_seeker
	a.Equal(expected, actual)
}

func setCreateSeekerRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/seeker_registration", s.Create_seeker(db))

	req, err := http.NewRequest(http.MethodPost, "/seeker_registration", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestCreateProviderAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_provider := m.Provider{
		Name:     "Spa",
		Email:    "",
		Password: "",
		Address:  "",
	}

	reqBody, err := json.Marshal(mock_provider)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateProviderRouter(db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Provider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.ProviderId = 0
	expected := mock_provider
	a.Equal(expected, actual)
}

func setCreateProviderRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/provider_registration", s.Create_service(db))

	req, err := http.NewRequest(http.MethodPost, "/provider_registration", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestProviderLoginAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_provider_login := m.Login{

		Email:    "spa@lakme.com",
		Password: "lakmepassword",
	}

	reqBody, err := json.Marshal(mock_provider_login)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setProviderLoginRouter(db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}
	t.Logf("Result %s", body)

	actual := m.Login{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.Login{}
	a.Equal(expected, actual)
}

func setProviderLoginRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/provider_login", s.Create_service(db))

	req, err := http.NewRequest(http.MethodPost, "/provider_login", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestSeekerLoginAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_seeker_login := m.Login{

		Email:    "nancy123@gmail.com",
		Password: "123",
	}

	reqBody, err := json.Marshal(mock_seeker_login)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setSeekerLoginRouter(db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Login{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.Login{}
	a.Equal(expected, actual)
}

func setSeekerLoginRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/seeker_login", s.Create_service(db))

	req, err := http.NewRequest(http.MethodPost, "/seeker_login", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestServicesAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	req, w, err := setServicesRouter(db)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Provider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.Provider{}
	a.Equal(expected, actual)
}

func setServicesRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.GET("/services", s.Listing_providers(db))

	req, err := http.NewRequest(http.MethodGet, "/services", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestServicesIDAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// service, err := insertServiceProvider(db)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setServicesIDRouter(db, "/:0")

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Provider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.Provider{}
	a.Equal(expected, actual)
}

func setServicesIDRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.GET("/services"+url, s.Listing_providers(db))

	req, err := http.NewRequest(http.MethodGet, "/services"+url, nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func insertServiceProvider(db *gorm.DB) (m.Provider, error) {
	s := m.Provider{
		ProviderId: 99,
		Name:       "test",
		Email:      "test@gmail.com",
		Password:   "test123",
	}

	if err := db.Create(&s).Error; err != nil {
		return s, err
	}

	return s, nil
}

func TestServiceRatingAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_rating := m.Ratings{
		Rating: 4,
	}

	reqBody, err := json.Marshal(mock_rating)
	if err != nil {
		a.Error(err)
	}
	end_url := "/:" + strconv.Itoa(int(mock_rating.ServiceID)) + "/rate_service"

	req, w, err := setServiceRateRouter(db, end_url, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Ratings{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.ServiceID = 0
	actual.ServiceName = ""
	expected := mock_rating

	a.Equal(expected, actual)
}

func setServiceRatingRouter(db *gorm.DB, url string, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/services"+url, s.Rate(db))

	req, err := http.NewRequest(http.MethodPost, "/services"+url, body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestServiceBookAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_service_book := m.Booking{
		ProviderId: 0,
		ServiceId:  1,
		SeekerName: "Lahari",
	}

	reqBody, err := json.Marshal(mock_service_book)
	if err != nil {
		a.Error(err)
	}
	end_url := "/:" + strconv.Itoa(int(mock_service_book.ServiceId)) + "/book"

	req, w, err := setServiceBookRouter(db, end_url, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Booking{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := mock_service_book
	expected.SeekerName = ""
	expected.ServiceId = 0
	a.Equal(expected, actual)
}

func setServiceBookRouter(db *gorm.DB, url string, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.GET("/services"+url, s.Book(db))

	req, err := http.NewRequest(http.MethodGet, "/services"+url, body)
	req.Header.Set("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImxhaGFyaUBnbWFpbC5jb20iLCJleHAiOjE2NDg5NTEzMDl9.8CCPJsoviPFjvp2ORrzKX1Hfl-PzUo0HzrBt0j6tcXM")
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestServiceRateAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_rating := m.Ratings{
		Rating: 4,
	}

	reqBody, err := json.Marshal(mock_rating)
	if err != nil {
		a.Error(err)
	}
	end_url := "/:" + strconv.Itoa(int(mock_rating.ServiceID)) + "/rate_service"

	req, w, err := setServiceRateRouter(db, end_url, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Ratings{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.ServiceID = 0
	actual.ServiceName = ""
	expected := mock_rating

	a.Equal(expected, actual)
}

func setServiceRateRouter(db *gorm.DB, url string, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/services"+url, s.Rate(db))

	req, err := http.NewRequest(http.MethodPost, "/services"+url, body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}
