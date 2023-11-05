package demo

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// target[start:end:max]
	// 切片从[2 3 4] 下标最大到5 cap(s) 5-1
	// 此时的第三个参数的意义low high max， low<=high<=max
	s := arr[1:4:5]
	fmt.Println(s, len(s), cap(s)) // [2 3 4] 3 4

}

func TestAppendSlice(t *testing.T) {
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d) %p\n", a1, len(a1), cap(a1), &a1)
	s9 := a1[1:4]

	fmt.Printf("s9: %v (len: %d, cap: %d) address %p \n", s9, len(s9), cap(s9), s9)
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d) address %p \n", i, s9, len(s9), cap(s9), s9)
	}
	fmt.Printf("a1: %v (len: %d, cap: %d) %p\n", a1, len(a1), cap(a1), &a1)

}
