package gmp

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNumCPU(t *testing.T) {
	// NumCPU 返回当前进程可以用到的逻辑核心数
	fmt.Println(runtime.NumCPU()) // mac 四核 ==> 逻辑核心数是8
}
