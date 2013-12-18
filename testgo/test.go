package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// f := closure(10)
	// fmt.Println(f(1))
	// fmt.Println(f(2))

	// fmt.Println(Test_defer())
	fmt.Println(Test_defer1())
	fmt.Println(Test_defer2())
}

func Test_defer() string {
	fmt.Println("a")
	defer fmt.Println("b")

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("_" + strconv.Itoa(i))
		}()
	}

	var result string

	defer func() {
		result = "test"
	}()

	fmt.Println("d")

	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		// 输出 2, 1, 0
		defer fmt.Println(i)
	}

	return result
}

// 最后结果返回 1， 说明内部匿名函数在return语句之后执行
func Test_defer1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func Test_defer2() (result int) {
	t := 5
	defer func() {
		t = t + 2
	}()
	return t
}

func Test_strings() {
	s := "sdfsdfsdf"
	s = s + "半"
	fmt.Println("%s", s)

	fmt.Println("strconv.FormatFloat : " + strconv.FormatFloat(10.100, 'f', -1, 32))

	fmt.Println("strings.Join : " + strings.Join([]string{"a", "b", "c"}, ","))

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //Fields are: ["foo" "bar" "baz"]
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

// 闭包
func closure(x int) func(int) int {
	fmt.Println("%p\n", &x)
	return func(y int) int {
		fmt.Println("%p\n", &x)
		return x + y
	}
}
