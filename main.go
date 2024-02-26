package main

import (
	"nuitee-task/app"
	"nuitee-task/config"
	"nuitee-task/exchangerates"
	"nuitee-task/handlers"
	"nuitee-task/hotelrates"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sLog := logger.Sugar()
	cfg := config.Config{}
	cfg.Default()

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		sLog.Fatalf("yamlFile.Get err   #%v ", err)
	}

	// Init the config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		sLog.Fatalf("unmarshal: %v", err)
	}

	// Init the hotelRates client
	hotelRates := hotelrates.NewHotelbedsClient(cfg.Apitude.APIKey, cfg.Apitude.Secret, cfg.Apitude.APIUrl)

	// Init the exchangeRates client and start the rate clearing process
	exchangeRates := exchangerates.NewCoinbaseClient(cfg.Coinbase.BaseURL, cfg.Coinbase.RefreshInterval)
	go exchangeRates.ClearRates()

	// Init the app that will process the requests
	app := app.NewApp(hotelRates, exchangeRates, sLog)

	// Init the HTTP server
	http := handlers.NewHTTP(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			app.Shutdown()
			os.Exit(1)
		}
	}()

	http.InitRoutes()
	http.Run(cfg.HTTP.Port)
}
