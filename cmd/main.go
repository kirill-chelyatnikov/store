package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"store/internal"
	"store/internal/store"
)

func main() {
	logger := initLogger()
	cfg := internal.GetConfig(logger)

	storeHandler := store.NewHandler()
	storeHandler.Register()

	httpServerStart(logger, cfg)
}

func httpServerStart(logger *logrus.Logger, cfg *internal.Config) {
	logger.Infof("starting HTTP server at port %d", cfg.Listen.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Listen.Host, cfg.Listen.Port), nil)
	if err != nil {
		logger.Fatalf("HTTP server can't start! error: %s", err)
	}
}

func initLogger() *logrus.Logger {
	log := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	return log
}
