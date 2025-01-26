package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// 생성자 함수
func newDict() *dict {
	d := new(dict)
	d.data = map[int]string{}
	return d
}

func main() {
	// 빈 객체 할당
	p := person{}

	// 값 할당
	p.name = "Lee"
	p.age = 10

	// {Lee 10}
	fmt.Println(p)

	// 실제 객체의 주소를 보고 싶다면 Printf와 %p를 사용하자
	fmt.Printf("%p\n", &p)

	var p1 person
	p1 = person{
		"Bob",
		20,
	}

	p2 := person{
		name: "minuk",
		age:  26,
	}

	// {Bob 20}
	fmt.Println(p1)
	// {minuk 26}
	fmt.Println(p2)

	// new는 빈 객체를 만든 후 "포인터"를 반환한다.
	pointer := new(person)
	pointer.name = "minah"

	// &{minah 0}
	fmt.Println(pointer)

	dic := newDict()

	dic.data[1] = "a"

	fmt.Println(dic.data)

}
