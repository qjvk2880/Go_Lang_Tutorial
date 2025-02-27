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

}
