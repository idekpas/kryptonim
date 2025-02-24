package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/idekpas/kryptonim/api/openexchangerates"
	"github.com/idekpas/kryptonim/config"
	"resty.dev/v3"
)

type ExchangeRepository interface {
	GetRates(string) (map[string]float64, error)
}

type OpenExchangeRepository struct {
	baseURL string
}

func NewOpenExchangeRepository() *OpenExchangeRepository {
	return &OpenExchangeRepository{baseURL: getBaseURL()}
}

func (oer OpenExchangeRepository) GetRates(currencies string) (rates map[string]float64, err error) {
	client := resty.New()
	defer client.Close()

	resp, err := client.R().
		SetQueryParam("base", "USD").
		SetQueryParam("symbols", currencies).
		SetResult(&openexchangerates.Response{}).
		Get(oer.baseURL)
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
