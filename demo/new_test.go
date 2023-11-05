package demo

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	p1 := new(int)

	i := 0
	i++
	p1 = &i

	fmt.Println(*p1)

	arr := make([]int, 0)
	if arr == nil {
		fmt.Println("arr is nil")

	} else {
		fmt.Printf("arr is not nill --> %#v \n ", arr) // []int{0, 0, 0}
	}

}
