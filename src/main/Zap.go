package main

import (
	"time"

	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	url := "www.google.com.hk"
	defer logger.Sync()
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly-typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	fmt.Println()
	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely-typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	fmt.Println()
	sugar.Infof("Failed to fetch URL: %s", url)
}
