package main

import (
	"fmt"
	// "github.com/Unknwon/com"
	"strings"
)

const (
	A = "A"
)

func main() {

	Test_map()
}

func Test_map() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)

	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[strings.ToUpper(v)] = k
	}
	fmt.Println(m2)

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
