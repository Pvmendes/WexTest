package config

import (
	"log"

	"github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion on .env file")
	}

	return value
}

// Initilize this variable to access the env values
var EnvConfigs *envConfigs

// We will call this in main.go to load the env variables
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// struct to map env values
type envConfigs struct {
	FiscalData_Base_URL                       string `mapstructure:"FISCALDATA_BASE_URL"`
	FiscalData_Endpoint_Rates_Exchange        string `mapstructure:"FISCALDATA_ENDPOINT_RATES_EXCHANGE"`
	FiscalData_Endpoint_PathSearch_Fields     string `mapstructure:"FISCALDATA_ENDPOINT_PATH_SEARCH_FIELDS"`
	FiscalData_Endpoint_PathSearch_Filter     string `mapstructure:"FISCALDATA_ENDPOINT_PATH_SEARCH_FILTER"`
	FiscalData_Endpoint_PathSearch_Sort       string `mapstructure:"FISCALDATA_ENDPOINT_PATH_SEARCH_SORT"`
	FiscalData_Endpoint_PathSearch_Pagination string `mapstructure:"FISCALDATA_ENDPOINT_PATH_SEARCH_PAGINATION"`
}

// Call to load the variables from env
func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("..")

	// Tell viper the name of your file
	viper.SetConfigName("app")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
