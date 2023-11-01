package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	j int = 0
	k int = 1
)

const (
	one = iota + 1 //0+1
	two            // 2
	// three = iota //2
	three // 3
	four  // 4
)

func main() {
	var i int = 10
	fmt.Println(i)
	fmt.Println("var block: ", j, k)

	var f32 float32 = 2.3
	var f64 float64 = 3.1415926
	fmt.Println("float type:", f32, f64)

	var bf bool = false
	var bt bool = true
	fmt.Println("bool type:", bf, bt)

	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("string type:", s1, s2, s1+s2)

	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println("zero value:", zi, zf, zb, zs)

	// short declare
	i1 := 32
	ss1 := "test"
	fmt.Println("short define:", i1, ss1)

	pi := &i1
	fmt.Println("point:", pi, *pi)

	const PI = 3.1414926
	fmt.Println("const value:", PI)

	var b byte = 221
	fmt.Printf("alia type byte=>uint8: %d, %T\n", b, b)

	var r rune = 3444
	fmt.Printf("alia type rune=>int32: %d, %T\n", r, r)

	fmt.Println("iota const:", one, two, three, four)

	// string <=> num
	i2s := strconv.Itoa(i)
	fmt.Printf("i: %T %d, i2s: %T %s\n", i, i, i2s, i2s)

	s2f, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Printf("string => float64: %T %f\n", s2f, s2f)

	f2s := strconv.FormatFloat(3.1415926, 'f', 8, 64)
	fmt.Printf("float64 => string: %T %s\n", f2s, f2s)

	// strings lib
	fmt.Println("prefix check:", strings.HasPrefix(s1, "H"))
	fmt.Println("index check:", strings.Index(s1, "o"))
	fmt.Println("to Upper:", strings.ToUpper(s1))

	var c complex64 = 2 + 3i
	fmt.Println("complex64:", c)
}
