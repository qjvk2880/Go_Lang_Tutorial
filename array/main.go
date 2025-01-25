package main

import "fmt"

func main() {
	//example()
	//stackExample()

	var a [3]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	b := [...]int{1, 2, 3}
	_ = b

	// 파라미터가 인터페이스라서 힙으로 이동함
	// 왜냐면 인터페이스에 값을 저장하면
	// 생명주기를 넘어 존재할 가능성이 있으니까
	fmt.Println(a)

	// 그런데 print를 지우면 탈출 x
	_ = a
}

func example() []int {
	s := make([]int, 5) // 힙 할당
	for i := range s {
		s[i] = i
	}
	return s // 슬라이스가 외부로 탈출
}

func stackExample() {
	s := make([]int, 3) // 배열은 스택에 할당
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s) // 스택에서만 사용
}
