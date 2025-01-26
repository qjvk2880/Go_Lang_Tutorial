package testlib

import "fmt"

var pop map[string]string

func init() {
	pop = map[string]string{
		"adele": "hello",
		"john":  "all of me",
	}
}

func GetMusic(singer string) string {
	printValues()
	return pop[singer]
}

func printValues() {
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
