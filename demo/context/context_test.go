package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// context包 -- 控制超时
// 控制超时，相当于我们同时监听两个channel，一个正常业务的结束channe, Done()返回
func TestTimeoutExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-bsChan:
		fmt.Println("business end ")

	}
}

func slowBusiness() {
	time.Sleep(time.Second * 2)
}

// context包 --time.AfterFunc控制超时

func TestTimeoutTimeAfter(t *testing.T) {
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("timeout")
	})
	<-bsChan
	fmt.Println("business end")

	timer.Stop()
}
