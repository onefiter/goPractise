package cache

import (
	"fmt"
	"runtime"
)

func CaclLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}

	case bool, uint8, int8:
		n = 1
	case int32, uint32, float32:
		n = 4
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}

	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.Value", value))
	}

	return n
}
