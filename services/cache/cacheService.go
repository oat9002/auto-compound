package services

type CacheService interface {
	SetWithoutExpiry(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
}
