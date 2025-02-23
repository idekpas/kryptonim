package services

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/idekpas/kryptonim/api/openexchangerates"
	"github.com/idekpas/kryptonim/config"
	"resty.dev/v3"
)

type RatesService interface {
	GetExchangeRates([]string) ([]map[string]interface{}, error)
}

type DefaultRatesService struct {
	baseURL string
}

func NewDefaultRatesService() *DefaultRatesService {
	return &DefaultRatesService{baseURL: getBaseURL()}
}

func (rs DefaultRatesService) GetExchangeRates(currencies []string) ([]map[string]interface{}, error) {
	rates, err := rs.getRates(strings.Join(currencies, ","))
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

func (rs DefaultRatesService) getRates(currencies string) (rates map[string]float64, err error) {
	client := resty.New()
	defer client.Close()

	resp, err := client.R().
		SetQueryParam("base", "USD").
		SetQueryParam("symbols", currencies).
		SetResult(&openexchangerates.Response{}).
		Get(rs.baseURL)
	if err != nil || resp.StatusCode() != 200 {
		log.Println("error during call to openexchangerates API: ", err)
		return nil, errors.New("error during call to openexchangerates API")
	}

	rates = resp.Result().(*openexchangerates.Response).Rates
	return rates, nil
}

func getBaseURL() string {
	cfg := config.GetConfig()
	url := cfg.GetString("api.url")
	key := cfg.GetString("api.key")

	return fmt.Sprintf("%s?app_id=%s", url, key)
}
