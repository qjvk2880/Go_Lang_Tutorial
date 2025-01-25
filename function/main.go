package main

import "fmt"

func main() {
	str := "hello"

	fmt.Println(str)

	changeStringByValue(str)
	fmt.Println(str)

	changeStringByReference(&str)
	fmt.Println(str)

	sumTotal := sum(1, 2, 3, 4, 5)
	fmt.Println(sumTotal)

	sumNameTotal, sumNameCount := sumNamedReturnParameter(1, 2, 3, 4, 5)
	fmt.Println(sumNameTotal, sumNameCount)
}

func changeStringByValue(str string) {
	str = "good bye"
}

func changeStringByReference(str *string) {
	*str = "good bye"
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sumNamedReturnParameter(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
		count++
	}
	return
}
