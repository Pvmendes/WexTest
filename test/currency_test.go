package test

import (
	"my-transaction-app/pkg/config"
	"my-transaction-app/pkg/currency"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestConvertCurrency(t *testing.T) {
	// load env variables just once in here so can be use in any other place
	config.InitEnvConfigs()

	// Mock the external API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response here based on the request
	}))
	defer server.Close()

	// Create an instance of your converter with the mocked server's URL
	//client :=  api.NewClient(server.URL)
	converter := currency.NewConverter()

	// Define the date for testing
	testDate := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC).String() // Example date

	// Define test cases
	tests := []struct {
		name          string
		amount        float64
		targetCountry string
		date          string
		want          float64
		wantErr       bool
	}{
		// Define various test cases here
		{
			name:          "Valid Conversion",
			amount:        100,
			targetCountry: "Portugal",
			date:          testDate,
			want:          85,
			wantErr:       false,
		},
		{
			name:          "No Rate Available",
			amount:        100,
			targetCountry: "Portugal",
			date:          testDate,
			want:          0,
			wantErr:       true,
		},
		/*{
			name:          "Invalid Country Code",
			amount:        100,
			targetCountry: "XXX",
			date:          testDate,
			want:          0,
			wantErr:       true,
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converter.Convert(tt.amount, tt.targetCountry, tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.ConvertedAmount != tt.want {
				t.Errorf("ConvertCurrency() got = %v, want %v", got.ConvertedAmount, tt.want)
			}
		})
	}
}
