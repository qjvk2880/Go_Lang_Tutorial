package main

import "fmt"

type calculator func(int, int) int

func main() {

	// 익명 함수를 변수에 선언
	add := func(a int, b int) int {
		return a + b
	}

	result1 := calc(add, 1, 2)
	fmt.Println(result1)

	// 함수를 파라미터에 직접 정의한다.
	result2 := calc(func(a int, b int) int {
		return a - b
	}, 1, 2)
	fmt.Println(result2)

	// 하지만 단점이 있다. 계속 func(~~~~~) int 부분이 반복된다.
	// 이 때 type 키워드를 사용할 수 있다.
}

// 함수를 파라미터로 받는 부분의 타입을 calculator로 바꾼다.
func calc(f calculator, a int, b int) int {
	return f(a, b)
}
