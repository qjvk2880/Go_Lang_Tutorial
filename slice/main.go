package main

import "fmt"

func main() {
	a := make([]int, 5)
	_ = a

	fmt.Println(a)
	fmt.Println(len(a))

	// append할 땐 새로 리턴한다.
	a = append(a, 1)

	// slice 두 개 concat
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
}
