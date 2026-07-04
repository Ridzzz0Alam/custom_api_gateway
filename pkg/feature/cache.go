package feature

import (
	"time"

	"github.com/ArmaanKatyal/porta/pkg/config"

	"github.com/patrickmn/go-cache"
)

type CacheExpiration int

const (
	DefaultExpiration CacheExpiration = 0
	NoExpiration      CacheExpiration = -1
)

type CacheHandler struct {
	Enabled            bool `json:"enabled"`
	ExpirationInterval uint `json:"expirationInterval"`
	CleanupInterval    uint `json:"cleanupInterval"`
	cache              *cache.Cache
}

func NewCacheHdnler(conf *config.CacheSettings) *CacheHandler {
	// If 0, set to default values

	if conf.ExpirationInterval == 0 {
		conf.ExpirationInterval = 5
	}
	if conf.Cleanupinterval == 0 {
		conf.Cleanupinterval = 10
	}
	return &CacheHandler{
		Enabled:            conf.Enabled,
		ExpirationInterval: conf.ExpirationInterval,
		CleanupInterval:    conf.Cleanupinterval,
		cache: cache.New(time.Duration(conf.ExpirationInterval)*time.Second,
			time.Duration(conf.Cleanupinterval)*time.Second),
	}
}

func (c *CacheHandler) Get(key string) (any, bool) {
	return c.cache.Get(key)
}

func (c *CacheHandler) Set(key string, value any, exp CacheExpiration) {
	c.cache.Set(key, value, time.Duration(exp))
}

func (c *CacheHandler) IsEnabled() bool {
	return c.Enabled
}
