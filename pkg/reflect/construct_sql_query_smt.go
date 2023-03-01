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
