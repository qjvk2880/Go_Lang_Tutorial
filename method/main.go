package main

import "fmt"

type Rect struct {
	width, height int
}

// 값 복사가 일어난다. Value Receiver
func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) areaByReference() int {
	r.height++
	return r.width * r.height
}

func main() {
	rect := Rect{
		12,
		12,
	}

	fmt.Println(rect.area())

	rect2 := Rect{
		12,
		13,
	}
	fmt.Println(rect2.areaByReference())
	fmt.Println(rect2)
}
