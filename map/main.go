package main

import "fmt"

func main() {
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map 키 체크
	val, exists := tickers["MST"]
	_ = val
	if !exists {
		fmt.Println("No MSFT ticker")
	}

	// 문자열은 값 타입이다.
	// 하지만 만약
	// s := "str"
	// s1 := s
	// 라고 한다면 다른 메타데이터를 가지지만 "str"이라는 데이터 공유한다.

	// 맵 같은 경우 값 타입의 value가 존재하지 않으면 zero 값을 리턴한다.
	// string의 zero값은 ""이다.
	noData := tickers["sss"]
	fmt.Println(noData == "")

	for key, val := range tickers {
		fmt.Println(key, val)
	}
	// TODO : 키로 정렬, Value로 정렬하는 법

}
