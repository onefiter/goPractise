package string

import (
	"fmt"
	"testing"
)

func TestReverseZhString(t *testing.T) {
	s := "abcdef你好"

	strArr := []rune(s)

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}

	fmt.Println(string(strArr))
}
