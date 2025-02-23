package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/idekpas/kryptonim/config"
)

type ExchangeService interface {
	Exchange(string, string, string) (map[string]interface{}, error)
}

type DefaultExchangeService struct {
	cryptoRates map[string]CryptoRate
}

type CryptoRate struct {
	DecimalPlaces int     `mapstructure:"decimalPlaces"`
	RateToUSD     float64 `mapstructure:"rateToUSD"`
}

func NewDefaultExchangeService() *DefaultExchangeService {
	cfg := config.GetConfig()
	var rates map[string]CryptoRate
	if err := cfg.UnmarshalKey("cryptoRates", &rates); err != nil {
		log.Fatalf("Error during mapping crypto rates from config: %v", err)
	}

	for k, v := range rates {
		fmt.Println("m1: ", k, " 2: ", v.DecimalPlaces, " 3: ", v.RateToUSD)
	}
	return &DefaultExchangeService{cryptoRates: rates}
}

func (des DefaultExchangeService) Exchange(from string, to string, amountStr string) (map[string]interface{}, error) {
	fromData, fromExists := des.cryptoRates[strings.ToLower(from)]
	toData, toExists := des.cryptoRates[strings.ToLower(to)]
	if !fromExists {
		log.Println("error exchanging crypto, currency does not exists11", fromData)
		return nil, errors.New("error exchanging crypto, currency does not exists")
	}

	if !toExists {
		log.Println("error exchanging crypto, currency does not exists22")
		return nil, errors.New("error exchanging crypto, currency does not exists")
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		log.Printf("error parsing amount: %v", err)
		return nil, err
	}

	convertedAmount := convert(amount, fromData.RateToUSD, toData.RateToUSD)
	roundedAmount := round(toData.DecimalPlaces, convertedAmount)

	return map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": roundedAmount,
	}, nil
}

func round(decimalPlaces int, convertedAmount float64) float64 {
	precision := math.Pow(10, float64(decimalPlaces))
	roundedAmount := math.Round(convertedAmount*precision) / precision
	return roundedAmount
}

func convert(amount float64, fromRate float64, toRate float64) float64 {
	amountInUSD := amount * fromRate
	convertedAmount := amountInUSD / toRate
	return convertedAmount
}
