package channel

import (
	"fmt"
	"testing"
)

// 从一个有缓冲的 channel 里读数据，当 channel 被关闭，依然能读出有效值。
// 只有当返回的 ok 为 false 时，读出的数据才是无效的。
func TestReadFromCloseChannel(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 18
	close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}
