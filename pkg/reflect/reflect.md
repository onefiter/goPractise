# reflect反射
> 反射是程序在运行时访问、检测和修改它本身状态或行为的一种能力，各种编程语言所实现的反射机制各有不同。

Go语言的`interface{}`类型变量具有解析任意类型变量的类型信息`type`和值信息`value`的能力。

Go的反射本质上就是利用interface{}的这种能力在运行时对任意变量的类型和值信息进行检视甚至是对值进行修改的机制。


## Go反射的三大原则

反射让静态类型语言Go在运行时具备了某种基于类型信息的动态特性。利用这种特性，fmt.Println在无法提前获知传入参数的真正类型的情况下依旧可以对其进行正确的格式化输出；json.Marshal也是通过这种特性对传入的任意结构体类型进行解构并正确生成对应的JSON文本。下面通过一个简单的构建SQL查询语句的例子来直观感受Go反射的“魔法”：

```go
package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func ConstructQueryStmt(obj interface{}) (stmt string, err error) {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		err = errors.New("only struct is supported")

		return
	}

	buffer := bytes.NewBufferString("")
	buffer.WriteString("SELECT ")

	if typ.NumField() == 0 {
		err = fmt.Errorf("the type[%s] has no fields", typ.Name())
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if i != 0 {
			buffer.WriteString(", ")
		}

		column := field.Name
		if tag := field.Tag.Get("orm"); tag != "" {
			column = tag
		}
		buffer.WriteString(column)
	}

	stmt = fmt.Sprintf("%s FROM %s", buffer.String(), strings.ToLower(typ.Name()))

	return
}

type Product struct {
	ID        uint32    `orm:"id"`
	Name      string    `orm:"name"`
	Price     uint32    `orm:"price"`
	LeftCount uint32    `orm:"left_count"`
	Batch     string    `orm:"batch_number"`
	Updated   time.Time `orm:"updated"`
}

type Person struct {
	ID      string    `orm:"id"`
	Name    string    `orm:"name"`
	Age     uint32    `orm:"age"`
	Gender  string    `orm:"gender"`
	Addr    string    `orm:"address"`
	Updated time.Time `orm:"updated"`
}

func main() {
	stmt, err := ConstructQueryStmt(&Product{})
	if err != nil {
		fmt.Println("construct query stmt for Product error: ", err)
		return
	}

	fmt.Println(stmt)

	stmt, err = ConstructQueryStmt(Person{})

	if err != nil {
		fmt.Println("construct query stmt for Person error: ", err)
	}

	fmt.Println(stmt)
}

```

```shell
go run main.go

SELECT id, name, price, left_count, batch_number, updated FROM product
SELECT id, name, age, gender, address, updated FROM person
```
Go反射十分适合处理这一类问题，它们的典型特点包括：

- 输入参数的类型无法提前确定；
- 函数或方法的处理结果因传入参数（的类型信息和值信息）的不同而异

---

反射的缺点
- 代码逻辑看起来不那么清晰，不容易理解
- 反射会让代码运行的更慢
- 在编译阶段，编译器无法检测到使用反射的代码中的问题（这种问题只能在Go程序运行时暴露出来，并且一旦暴露，很大可能会导致运行时的panic）。


>Rob Pike为Go反射的规范使用定义了三大法则，如果经过评估，你必须使用反射才能实现你要的功能特性，那么你在使用反射时需要牢记这三条法则。
1. 反射世界的入口：经由接口（`interface{}`）类型变量值进入反射的世界并获得对应的反射对象（`reflect.Value或reflect.Type`）。
2. 反射世界的出口：反射对象（`reflect.Value`）通过化身为一个接口（`interface{}`）类型变量值的形式走出反射世界。
3. 修改反射对象的前提：反射对象对应的reflect.Value必须是可设置的（`Settable`）。