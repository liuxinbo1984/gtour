package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// testArr()
	// testMap()

	testString2Byte()
}

func testString2Byte() {
	s := "Hello 雪山飞狐"
	bs := []byte(s)
	fmt.Println(bs[1], bs[5], bs[15])

	fmt.Println("byte len:", len(s))
	fmt.Println("Charcter/Unicode len:", utf8.RuneCountInString(s))

	for i, v := range s {
		// fmt.Println(i, "=", v, s[i])
		fmt.Printf("index:%d, character:%c, unicode:%d\n", i, v, v)
	}
}

func testMap() {

	map1 := map[string]int{}
	map1["arvin"] = 31
	map1["shadow"] = 26

	fmt.Println(map1["arvin"])
	// key不存在，会返回零值： 0
	fmt.Println(map1["nobody"])
	v, ok := map1["shadow"]
	if ok {
		// 检查是否存在
		fmt.Println(v)
	}

	delete(map1, "shadow")
	fmt.Println(map1)

	// 遍历map
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

	// 元素数量
	fmt.Println(len(map1))
}

func testArr() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr[2])
	fmt.Println(arr)

	// 自动计算长度
	arr2 := [...]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr2)

	// 按索引位赋值， 长度以【声明长度】为准
	arr3 := [5]string{1: "b", 3: "d"}
	fmt.Println(arr3)

	for i, v := range arr3 {
		fmt.Printf("index: %d, value: |%-2s|\n", i, v)
	}

	// slice：底层引用数组
	var s1 []string = arr[2:5]
	fmt.Printf("slice1 %+v\n", s1)

	s2 := []string{"a", "b", "c"}
	fmt.Printf("slice2: len %d, cap %d\n", len(s2), cap(s2))

	s3 := make([]string, 4, 8)
	fmt.Printf("old s3 address:%p\n", &s3)
	fmt.Printf("slice3: len %d, cap %d\n", len(s3), cap(s3))

	s3 = append(s3, "e", "f", "g")
	fmt.Printf("s3 address:%p\n", &s3)
	s3 = append(s3, s2...)
	fmt.Printf("s3 address:%p\n", &s3)
	fmt.Println(s3)
}
