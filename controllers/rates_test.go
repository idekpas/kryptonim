package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockRatesService struct{}

func newMockRatesService() *mockRatesService {
	return &mockRatesService{}
}

func (mrs mockRatesService) GetExchangeRates(currencies []string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{"from": "USD", "to": "EUR", "rate": 0.85},
		{"from": "EUR", "to": "USD", "rate": 1.18},
	}, nil
}

func TestGetRates_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := RatesController{RatesService: newMockRatesService()}
	router.GET("/rates", controller.GetRates)

	req, _ := http.NewRequest("GET", "/rates?currencies=USD,EUR", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	expectedBody := `[{"from":"USD","to":"EUR","rate":0.85},{"from":"EUR","to":"USD","rate":1.18}]`
	assert.JSONEq(t, expectedBody, resp.Body.String())
}

func TestGetRates_BadRequest_NoCurrencies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := RatesController{}
	router.GET("/rates", controller.GetRates)

	req, _ := http.NewRequest("GET", "/rates?currencies=", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.JSONEq(t, `{}`, resp.Body.String())
}

func TestGetRates_BadRequest_OneCurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := RatesController{}
	router.GET("/rates", controller.GetRates)

	req, _ := http.NewRequest("GET", "/rates?currencies=USD", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.JSONEq(t, `{}`, resp.Body.String())
}
