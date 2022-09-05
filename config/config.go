package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type HullCityConfig struct {
	BaseUrl   string
	FetchPath string
	GetIdPath string
}

type Config struct {
	DatabaseURL          string
	DatabaseName         string
	CronFrequencyMinutes int
	QuantityToFetch      int
	HttpPort             string
	HullCityConfig       HullCityConfig
}

// LoadConfig will load config from environment variable.
func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	qty, err := strconv.Atoi(os.Getenv("QUANTITY_TO_FETCH"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	cronFreq, err := strconv.Atoi(os.Getenv("CRON_FREQUENCY_MINUTES"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return &Config{
		DatabaseURL:          os.Getenv("DATABASE_URL"),
		DatabaseName:         os.Getenv("DATABASE_NAME"),
		QuantityToFetch:      qty,
		CronFrequencyMinutes: cronFreq,
		HttpPort:             os.Getenv("HTTP_PORT"),
		HullCityConfig: HullCityConfig{
			BaseUrl:   os.Getenv("HULL_CITY_BASE_URL"),
			FetchPath: os.Getenv("HULL_CITY_FETCH_PATH"),
			GetIdPath: os.Getenv("HULL_CITY_GET_PATH"),
		},
	}
}
