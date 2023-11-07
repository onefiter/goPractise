package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ************ 简单原生类型 *****************
	var b = true
	val := reflect.ValueOf(b)
	typ := reflect.TypeOf(b)
	// bool true
	fmt.Println(typ.Name(), val.Bool())

	var i = 23 // 整型
	val = reflect.ValueOf(i)
	typ = reflect.TypeOf(i)
	// int 23
	fmt.Println(typ.Name(), val.Int())

	var f = 3.14 // 浮点型
	val = reflect.ValueOf(f)
	typ = reflect.TypeOf(f)
	// float64 3.14
	fmt.Println(typ.Name(), val.Float())

	var s = "hello, reflection" // 字符串
	val = reflect.ValueOf(s)
	typ = reflect.TypeOf(s)
	// string hello, reflection
	fmt.Println(typ.Name(), val.String())

	var fn = func(a, b int) int { // 函数（一等公民）
		return a + b
	}
	val = reflect.ValueOf(fn)
	typ = reflect.TypeOf(fn)
	// func func(int, int) int
	// 函数的typ.Name() 返回的是空
	fmt.Println(typ.Kind(), typ.String())

	// 指针
	var pi01 = (*int)(nil)
	var ps = (*string)(nil)
	typ = reflect.TypeOf(pi01)
	// prt *int
	fmt.Println(typ.Kind(), typ.String())

	typ = reflect.TypeOf(ps)
	// ptr *string
	fmt.Println(typ.Kind(), typ.String())

	//*******************原生符合类型**************************
	var s1 = []int{5, 6} // 切片
	val = reflect.ValueOf(s1)
	typ = reflect.TypeOf(s1)
	// [5,6]
	fmt.Printf("[%d %d]\n", val.Index(0).Int(), val.Index(1).Int())
	// slice []int
	fmt.Println(typ.Kind(), typ.String())

	// 数组
	var arr = [3]int{5, 6}
	val = reflect.ValueOf(arr)
	typ = reflect.TypeOf(arr)
	// [5 6 0]
	fmt.Printf("[%d %d %d]\n", val.Index(0).Int(), val.Index(1).Int(), val.Index(2).Int())
	// array [3]int
	fmt.Println(typ.Kind(), typ.String())

	// map
	var m = map[string]int{
		"tony": 1,
		"jim":  2,
		"johb": 3,
	}
	val = reflect.ValueOf(m)
	typ = reflect.TypeOf(m)
	iter := val.MapRange()
	fmt.Printf("{")
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("%s:%d,", k.String(), v.Int())
	}
	// {tony:1,jim:2,john:3,}
	fmt.Printf("}\n")
	// map map[string]int
	fmt.Println(typ.Kind(), typ.String())

	// 结构体
	type Person struct {
		Name string
		Age  int
	}
	var p = Person{"tony", 3}
	val = reflect.ValueOf(p)
	typ = reflect.TypeOf(p)
	// {tony 3}
	fmt.Printf("{%s %d}\n", val.Field(0).String(), val.Field(1).Int())
	// struct Person main.Person
	fmt.Println(typ.Kind(), typ.Name(), typ.String())

	// channel
	var ch = make(chan int, 1)
	val = reflect.ValueOf(ch)
	typ = reflect.TypeOf(ch)
	ch <- 17
	v, ok := val.TryRecv()
	if ok {
		// 17
		fmt.Println(v.Int())
	}
	// channel channel int
	fmt.Println(typ.Kind(), typ.String())

	// 其他自定义类型
	type MyInt int

	var mi MyInt = 19
	val = reflect.ValueOf(mi)
	typ = reflect.TypeOf(mi)
	// MyInt int main.MyInt 19
	fmt.Println(typ.Name(), typ.Kind(), typ.String(), val.Int())

}
