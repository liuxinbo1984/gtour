package main

import (
	"fmt"
)

func main() {
	// testIfAndSwitch()
	testFor()
}

func testFor() {
	sum := 0
	for i := 0; i <= 100; i++ {
		// i 块级作用域
		sum += i
	}
	fmt.Println("for1:", sum)

	sum1, i := 0, 0
	for i <= 100 {
		sum1 += i
		i++
	}
	fmt.Println("for2:", sum1)

	sum2, i := 0, 0
	for {
		sum2 += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println("for3:", sum2)

	// go中没有while
}

func testIfAndSwitch() {
	i := 6
	if i > 10 {
		fmt.Println(">10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<5")
	}

	// 注意【逗号】，有初始化,初始化变量的作用域限定在switch块中
	switch i := 7; {
	case i > 10:
		{
			fmt.Println(">10")
		}
	case i > 5 && i <= 10:
		{
			fmt.Println("switch1: 5<i<=10", i)
		}
	default: //默认
		{
			fmt.Println("i<5")
		}
	}

	// 没有初始化，直接判断
	switch {
	case i > 10:
		{
			fmt.Println(">10")
		}
	case i > 5 && i <= 10:
		{
			fmt.Println("switch2: 5<i<=10", i)
		}
		// 不用加break，匹配即中断
		// 如果需要匹配下一个分支，使用 fallthrough
		// fallthrough
	default: //默认
		{
			fmt.Println("i<5")
		}
	}

	switch i := 1; i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没匹配！")
	}
}
