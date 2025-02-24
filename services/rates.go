package services

import (
	"errors"
	"log"
	"strings"

	"github.com/idekpas/kryptonim/repository"
)

type RatesService interface {
	GetExchangeRates([]string) ([]map[string]interface{}, error)
}

type DefaultRatesService struct {
	rep repository.ExchangeRepository
}

func NewDefaultRatesService() *DefaultRatesService {
	return &DefaultRatesService{rep: repository.NewOpenExchangeRepository()}
}

func (rs DefaultRatesService) GetExchangeRates(currencies []string) ([]map[string]interface{}, error) {
	rates, err := rs.rep.GetRates(strings.Join(currencies, ","))
	if err != nil || len(rates) < 2 {
		log.Println("error fetching rates from response: ", err)
		return nil, errors.New("error fetching rates from response")
	}

	resp := calculateRates(currencies, rates)

	return resp, nil
}

func calculateRates(currencies []string, rates map[string]float64) []map[string]interface{} {
	var resp []map[string]interface{}
	for i, from := range currencies {
		for j, to := range currencies {
			if i != j {
				resp = append(resp, map[string]interface{}{
					"from": from,
					"to":   to,
					"rate": rates[to] / rates[from],
				})
			}
		}
	}
	return resp
}
