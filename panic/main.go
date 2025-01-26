package main

import (
	"fmt"
	"os"
)

func main() {
	// 잘못된 파일명을 넣음
	openFile("Invalid.txt")

	// recover에 의해
	// 이 문장 실행됨
	println("Done")
}

func openFile(fn string) {
	// defer 함수. panic 호출시 실행됨
	var r interface{}
	defer func() {
		// r을 수행한다.
		// 그 후 정상적으로 리커버가 수행 되었다면
		// OPEN ERROR open Invalid.txt: no such file or directory
		// r에는 에러 메세지가 있다.

		// 그리고 패닉이 실행되면 defer 함수들을 호출하므로 recover는 꼭 defer 함수 내부에서 사용한다.
		if r = recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}

		fmt.Println(r)
	}()
	// 위에는 function "call"이 필요해서 소괄호

	f, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer f.Close()
}
