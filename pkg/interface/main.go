package main

import (
	"fmt"
)

type Car interface {
	Drive()
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
	fmt.Println(c)
	
}