package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExchangeRepository struct {
	mock.Mock
}

func (m *MockExchangeRepository) GetRates(currencies string) (map[string]float64, error) {
	args := m.Called(currencies)
	return args.Get(0).(map[string]float64), args.Error(1)
}

func TestGetExchangeRates_Success(t *testing.T) {
	mockRepo := new(MockExchangeRepository)
	mockRepo.On("GetRates", "USD,EUR,GBP").Return(map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"GBP": 0.75,
	}, nil)

	service := DefaultRatesService{rep: mockRepo}
	rates, err := service.GetExchangeRates([]string{"USD", "EUR", "GBP"})

	assert.NoError(t, err)
	assert.Len(t, rates, 6)
}

func TestGetExchangeRates_InvalidCurrencies(t *testing.T) {
	mockRepo := new(MockExchangeRepository)
	mockRepo.On("GetRates", "USD").Return(map[string]float64{"USD": 1.0}, nil)

	service := DefaultRatesService{rep: mockRepo}
	rates, err := service.GetExchangeRates([]string{"USD"})

	assert.Error(t, err)
	assert.Nil(t, rates)
}

func TestCalculateRates(t *testing.T) {
	currencies := []string{"USD", "EUR", "GBP"}
	rates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"GBP": 0.75,
	}

	expected := []map[string]interface{}{
		{"from": "USD", "to": "EUR", "rate": 0.85},
		{"from": "USD", "to": "GBP", "rate": 0.75},
		{"from": "EUR", "to": "USD", "rate": 1.1764705882352942},
		{"from": "EUR", "to": "GBP", "rate": 0.8823529411764706},
		{"from": "GBP", "to": "USD", "rate": 1.3333333333333333},
		{"from": "GBP", "to": "EUR", "rate": 1.1333333333333333},
	}

	result := calculateRates(currencies, rates)

	assert.Equal(t, expected, result)
}
