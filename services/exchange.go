package services

import (
	"errors"
	"log"
	"math"
	"math/big"
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

	return &DefaultExchangeService{cryptoRates: rates}
}

func (des DefaultExchangeService) Exchange(from string, to string, amountStr string) (map[string]interface{}, error) {
	fromData, fromExists := des.cryptoRates[strings.ToLower(from)]
	toData, toExists := des.cryptoRates[strings.ToLower(to)]
	if !fromExists || !toExists {
		log.Println("error exchanging crypto, currency does not exists")
		return nil, errors.New("error exchanging crypto, currency does not exists")
	}

	fromRateBig := floatToBigInt(fromData.RateToUSD, fromData.DecimalPlaces)
	toRateBig := floatToBigInt(toData.RateToUSD, toData.DecimalPlaces)

	amountBig := new(big.Int)
	amountBig.SetString(amountStr, 10)

	amountInUSD := new(big.Int).Mul(amountBig, fromRateBig)
	convertedAmount := new(big.Int).Div(amountInUSD, toRateBig)

	return map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": bigIntToFloat64(convertedAmount, toData.DecimalPlaces),
	}, nil
}

func floatToBigInt(value float64, decimalPlaces int) *big.Int {
	multiplier := math.Pow10(decimalPlaces)
	temp := big.NewFloat(value * multiplier)
	result := new(big.Int)
	temp.Int(result)
	return result
}

func bigIntToFloat64(value *big.Int, decimalPlaces int) float64 {
	divisor := math.Pow10(decimalPlaces)
	temp := new(big.Float).SetInt(value)
	result, _ := temp.Quo(temp, big.NewFloat(divisor)).Float64()
	return result
}
