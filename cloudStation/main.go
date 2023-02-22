package main

import (
	"fmt"

	"github.com/goPractise/cloudStation/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
