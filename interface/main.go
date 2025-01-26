package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 인터페이스의 모든 메소드를 구현해야한다.
// 그렇지 않으면 컴파일 에러가 발생한다.
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func main() {
	r := Rect{10., 20.}

	c := Circle{radius: 10}

	showArea(r, c)

	// 빈 인터페이스는, 모든 타입이 적어도 0개의 메소드를 구현하므로 모든 타입을 나타낸다.
	// 즉, 자바의 객체 최고 조상인 Object와 비슷한 개념같다.
	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	println(i) // 포인터주소 출력
	println(j)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}
