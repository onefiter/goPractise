package demo

import (
	"fmt"
	"testing"
)

func TestCopySlice(t *testing.T) {
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println("refData[0] = ", refData[0]) // 999
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	fmt.Println("srcData[4:6] = ", srcData[4:6])
	// srcData[4:6] => [4 5]
	// 将[4 5]拷贝到copyData中
	copy(copyData, srcData[4:6])
	for i := 0; i < elementCount-1; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}

func TestCopySliceV2(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println("slice1 to slice2 ", slice1, slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println("slice2 to slice1 ", slice1, slice2)

}
