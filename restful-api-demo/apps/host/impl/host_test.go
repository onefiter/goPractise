package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/goPractise/restful-api-demo/apps/host"
	"github.com/goPractise/restful-api-demo/apps/host/impl"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	// 需要初始化全局Logger,
	// 为什么不设计为默认打印, 因为性能
	fmt.Println(zap.DevelopmentSetup())

	// host service 的具体实现
	service = impl.NewHostServiceImpl()
}
