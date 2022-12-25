package lru_test

import (
	"testing"

	"github.com/goPractise/cache/lru"
	"github.com/matryer/is"
)

func TestOnEvicted(t *testing.T) {
	is := is.New(t)

	keys := make([]string, 0, 8)
	onEvicted := func(key string, value interface{}) {
		keys = append(keys, key)
	}

	cache := lru.New(16, onEvicted)

	cache.Set("k1", 1)
	cache.Set("k2", 2)
	cache.Get("k1")
	cache.Set("k3", 3)
	cache.Get("k1")
	cache.Set("k4", 4)

	expected := []string{"k2", "k3"}
	is.Equal(expected, keys)
	is.Equal(2, cache.Len())
}
