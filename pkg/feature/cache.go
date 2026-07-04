package feature

import (
	"time"

	"github.com/ArmaanKatyal/porta/pkg/config"

	"github.com/patrickmn/go-cache"
)

type CacheExpiration interface

const (
	DefaultExpiration CacheExpiration = 0
	NoExpiration CacheExpiration = -1 
)

type CacheHandler struct {
	Enabled				bool `json:"enabled"`
	ExpirationInterval	uint
}