package cache

import "time"

type CacheService interface {
	SetWithoutExpiry(key string, value interface{})
	Set(key string, value interface{}, expiry time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
}
