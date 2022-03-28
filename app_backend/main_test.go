package main

import (
	s "app_backend/controllers"
	m "app_backend/model"

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

func TestCreateServiceAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock_provider := m.ServiceAndProvider{
		ServiceName:        "Spa",
		ProviderEmail:      "spa@lakme.com",
		ProviderPassword:   "lakmepassword",
		ServicePrice:       500,
		ServiceDescription: "Spa at home in 120 mins",
	}

	reqBody, err := json.Marshal(mock_provider)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateServiceRouter(db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.ServiceAndProvider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := mock_provider
	a.Equal(expected, actual)
}

func setCreateServiceRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/service_registration", s.Create_service(db))

	req, err := http.NewRequest(http.MethodPost, "/service_registration", body)
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

	actual := m.ServiceAndProvider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.ServiceAndProvider{}
	a.Equal(expected, actual)
}

func setServicesRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.GET("/services", s.Listing_services(db))

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

	actual := m.ServiceAndProvider{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.ServiceAndProvider{}
	a.Equal(expected, actual)
}

func setServicesIDRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.GET("/services"+url, s.Listing_services(db))

	req, err := http.NewRequest(http.MethodGet, "/services"+url, nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func insertServiceProvider(db *gorm.DB) (m.ServiceAndProvider, error) {
	s := m.ServiceAndProvider{
		ServiceId:          99,
		ServiceName:        "test",
		ServicePrice:       100,
		ServiceDescription: "none",
		ProviderEmail:      "test@gmail.com",
		ProviderPassword:   "test123",
	}

	if err := db.Create(&s).Error; err != nil {
		return s, err
	}

	return s, nil
}

func TestServiceBookAPI(t *testing.T) {

	a := assert.New(t)

	db, err := gorm.Open("sqlite3", "./dbase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	req, w, err := setServiceBookRouter(db, "/:8/book")

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := m.Booking{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := m.Booking{}
	a.Equal(expected, actual)
}

func setServiceBookRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/services"+url, s.Book(db))

	req, err := http.NewRequest(http.MethodPost, "/services"+url, nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}
