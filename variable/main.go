package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	rawLiteral := `안녕하세요\n
                   여기의 \n은 해석되지 않는다.\n`
	fmt.Println(rawLiteral)

	interLiteral := "여기의 \n은 해석된다.\n"
	fmt.Println(interLiteral)
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//                 "아리리요"
	interLiteral = "여기의 \n" +
		"은 해석된다.\n"
	fmt.Println(interLiteral)

	// 암묵적 변환을 지원하지 않는다.
	// 따라서 명시적 변환을 해줘야 한다.

	// int8....int64까지 있다.
	// 일반적인 int는 64bit 아키텍처에서 int64, 32bit이라면 int32

	var i int = 100
	var f float32 = float32(i)

	fmt.Println(i, f)

	str := "ABC"
	bytes := []byte(str)

	// 배열은 [65, 66, 67] 형태로 출력된다.
	fmt.Println(bytes)
}
