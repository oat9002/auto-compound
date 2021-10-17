package services

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(expiry time.Duration, clearExpiredCacheInterval time.Duration) *CacheService {
	c := cache.New(expiry, clearExpiredCacheInterval)
	cacheService := &CacheService{cache: c}

	return cacheService
}

func (c *CacheService) SetWithoutExpiry(key string, value interface{}) {
	c.cache.Set(key, value, cache.NoExpiration)
}

func (c *CacheService) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}
