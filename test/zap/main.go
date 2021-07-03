package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction(zap.Development())
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "hello"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Infof("hello %s", url)
	log, _ := zap.NewDevelopment()
	log.Debug("hello world")
	log.Info("hello data", zap.String("url", url))
	slogger := log.Sugar()
	slogger.Infof("hello %s", url)

}
