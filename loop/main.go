package main

import "fmt"

func main() {
	sum := 0

	// if 문과 동일하게 소괄호가 없다.
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	n := 1
	// while 문처럼 사용이 가능하다.
	for n < 100 {
		n *= 2
		if n > 60 {
			break
		}
	}
	fmt.Println(n)

	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		fmt.Println(index, name)
	}
}
