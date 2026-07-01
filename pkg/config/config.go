package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/sony/gobreaker/v2"
	"github.com/spf13/viper"
)

func init() {
	serviceName := "api_gateway"

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("admin.port", "8081")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.metrics.prefix", serviceName)
	viper.SetDefault("server.metrics.buckets", []float64{0.005, 0.01, 0.025, 0.05, 0.1})
}

type CircuitSettings struct {
	Enabled			bool	`yaml:"enabled"`
	Timeout			uint	`yaml:"timeout"`
	Interval		uint	`yaml:"interval"`
	FailureRatio	float64	`yaml:"failureRatio"`
}