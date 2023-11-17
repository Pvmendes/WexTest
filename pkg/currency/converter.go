package currency

import (
	"errors"
	"fmt"
	"my-transaction-app/pkg/api"
	"my-transaction-app/pkg/utils"
	"strconv"
	"strings"
	"time"
)

// ConversionResult holds the result of a currency conversion.
type ConversionResult struct {
	ConvertedAmount  float64
	ExchangeRate     float64
	ExchangeRateDate string
}

type Converter struct {
	api *api.ClientTreasuryReporting
}

func NewConverter() *Converter {
	return &Converter{
		api: api.NewClient(),
	}
}

func (c *Converter) Convert(amount float64, toCountry string, date string) (*ConversionResult, error) {

	exchangeItem, err := c.api.GetExchangeRate(date, toCountry)
	if err != nil {
		return new(ConversionResult), err
	}

	// Perform conversion logic
	exchangeRate, _ := strconv.ParseFloat(strings.TrimSpace(exchangeItem.ExchangeRate), 64)
	convertedAmount := amount * exchangeRate
	fmt.Println(exchangeRate)

	recordDate, _ := time.Parse("2006-01-02T15:04:05.000Z", exchangeItem.RecordDate)
	purchaseDate, _ := time.Parse("2006-01-02T15:04:05.000Z", date)

	isValid := isRateDateValid(recordDate, purchaseDate)
	if !isValid {
		return new(ConversionResult), errors.New("exchange rate date " + exchangeItem.RecordDate + " is not valid for the purchase date " + date)
	}

	// Return converted amount or an error
	return &ConversionResult{
		ConvertedAmount:  utils.RoundFloat(convertedAmount, 2),
		ExchangeRate:     exchangeRate,
		ExchangeRateDate: exchangeItem.RecordDate,
	}, nil
}

// isRateDateValid checks if the rateDate is within 6 months before the purchaseDate and not after it.
func isRateDateValid(rateDate, purchaseDate time.Time) bool {
	sixMonthsBefore := purchaseDate.AddDate(0, -6, 0) // 6 months before the purchase date
	return !rateDate.After(purchaseDate) && !rateDate.Before(sixMonthsBefore)
}
