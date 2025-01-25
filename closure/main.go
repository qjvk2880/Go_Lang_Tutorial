package main

import (
	"fmt"
)

func nextValue() func() int {
	i := 0
	// 이 i 변수가 힙 영역으로 이동한다.
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()

	// 해당 next의 타입은 func() int이다.
	fmt.Println(next())
	fmt.Println(next())
}
