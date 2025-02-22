package cache

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

var cache *lru.Cache[string, string]

func Init(maxEntries int) {
	cache, _ = lru.New[string, string](maxEntries)
}

func Get(key string) (string, bool) {
	return cache.Get(key)
}

func Set(key, value string) {
	cache.Add(key, value)
}
