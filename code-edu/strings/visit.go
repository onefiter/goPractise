package main

import "fmt"

func main() {
	s := "斗破苍穹2023之onefiter"

	// s := "onefiter"
	fmt.Println("字符串长度:", len(s))
	// 安装字节数进行访问
	for i := 0; i < len(s); i++ {
		// 会打印字节的ASCII
		fmt.Println(s[i])

	}

	for _, char := range s {
		fmt.Printf("%c\n", char)
	}
}
