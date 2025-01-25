package main

import "fmt"

func main() {
	var k interface{}

	k = 1
	if k == 1 { // 조건문에 소괄호를 붙이지 않는다.
		fmt.Println("One")
	}

	switch v := k.(type) {
	case int:
		fmt.Println("정수")
		fmt.Println(v)
	case string:
		fmt.Println("문자열")
		fmt.Println(v)
	}

}
