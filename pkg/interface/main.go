package main

import (
	"fmt"
)

type Car interface {
	Drive()
}


type TrafficTool interface {
	Drive()
	Run()
}

type Truck  struct {
	Model string
}

func (t Truck) Drive() {
	fmt.Println(t.Model)
}

func main() {
	// c并不是简单地转为了Truck类型
	var c Car = Truck{}

	// // 断言作用1：类型转换
	// t := c.(Truck)
	// 此时t是Truck类型
	// fmt.Println(t)
	

	// 比较了底层的实现，和目标的接口是不是实现了目标接口的方法
	// 当TrafficTool中的接口c中并没有完全实现，就意味着不能转
	// tt := c.(TrafficTool)
	// fmt.Println(tt)

	switch c.(type) {
	case TrafficTool:
		fmt.Println("TrafficTool")
	case Truck:
		fmt.Println("Truck")
	case Car:
		fmt.Println("Car")
	default:
		fmt.Println("None")
		
	}
}