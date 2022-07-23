package services

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type InMemCacheService struct {
	cache *cache.Cache
}

func NewInMemCacheService(expiry time.Duration, clearExpiredCacheInterval time.Duration) *InMemCacheService {
	c := cache.New(expiry, clearExpiredCacheInterval)
	cacheService := &InMemCacheService{cache: c}

	return cacheService
}

func (c *InMemCacheService) SetWithoutExpiry(key string, value interface{}) {
	c.cache.Set(key, value, cache.NoExpiration)
}

func (c *InMemCacheService) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *InMemCacheService) Delete(key string) {
	c.cache.Delete(key)
}
