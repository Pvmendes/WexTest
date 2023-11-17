package api

import (
	"errors"
	"fmt"
	"my-transaction-app/pkg/config"
	"my-transaction-app/pkg/utils"
	"slices"
)

// Client is responsible for fetching the exchange rates
type ClientTreasuryReporting struct {
	baseURL               string
	endPointRatesExchange string
	fields                string
	filter                string
	sort                  string
	pagination            string
}

// ExchangeRateResponse represents the response structure from the API.
type ExchangeRateResponse struct {
	Data []ExchangeRateItemResponse `json:"data"`
}

type ExchangeRateItemResponse struct {
	RecordDate          string `json:"record_date"`
	Country             string `json:"country"`
	Currency            string `json:"currency"`
	CountryCurrencyDesc string `json:"country_currency_desc"`
	ExchangeRate        string `json:"exchange_rate"`
}

func NewClient() *ClientTreasuryReporting {
	return &ClientTreasuryReporting{
		baseURL:               config.EnvConfigs.FiscalData_Base_URL,
		endPointRatesExchange: config.EnvConfigs.FiscalData_Endpoint_Rates_Exchange,
		fields:                config.EnvConfigs.FiscalData_Endpoint_PathSearch_Fields,
		filter:                config.EnvConfigs.FiscalData_Endpoint_PathSearch_Filter,
		sort:                  config.EnvConfigs.FiscalData_Endpoint_PathSearch_Sort,
		pagination:            config.EnvConfigs.FiscalData_Endpoint_PathSearch_Pagination,
	}
}

// Call the Treasury Reporting Rates of Exchange API
func (c *ClientTreasuryReporting) GetExchangeRate(date string, country string) (ExchangeRateItemResponse, error) {
	exechangeResponse := &ExchangeRateResponse{}

	// Construct the URL with parameters
	finalUrl := c.baseURL + c.endPointRatesExchange + c.fields

	finalUrl += fmt.Sprintf(c.filter, date, country)

	finalUrl += c.sort

	finalUrl += c.pagination

	err := utils.GetJson(finalUrl, exechangeResponse)
	if err != nil {
		return ExchangeRateItemResponse{}, err
	}

	idx := slices.IndexFunc(exechangeResponse.Data, func(c ExchangeRateItemResponse) bool { return c.Country == country })
	if idx == -1 {
		return ExchangeRateItemResponse{}, errors.New("No item from API has found")
	}

	// Return the exchange rate or error
	return exechangeResponse.Data[idx], nil
}
