package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad error"),
}

func bad() bool {
	return false
}

func returnError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

// error: <nil>
func main() {
	e := returnError()
	if e != nil {
		fmt.Printf("error: %+v\n", e)
		return
	}

	fmt.Println("ok")
}
