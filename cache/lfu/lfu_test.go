package lfu_test

import (
	"testing"

	"github.com/goPractise/cache/lfu"
	"github.com/matryer/is"
)

func TestSet(t *testing.T) {
	is := is.New(t)

	cache := lfu.New(24, nil)
	cache.DelOldest()
	cache.Set("k1", 1)
	v := cache.Get("k1")

	is.Equal(v, 1)

	cache.Del("k1")
	is.Equal(0, cache.Len())

}
