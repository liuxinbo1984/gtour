package main

import (
	"errors"
	"fmt"
)

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	r1, err := sum1(-1, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1)
	}

	r2, _ := sum2(3, 4)
	fmt.Println(r2)

	fmt.Println("可变参数函数：", sum3(1, 2, 3, 4, 5))

	c := closure()
	fmt.Println("闭包：", c(), c(), c())

	// 方法
	a := Age(20)
	a.String()
	// (&a).Modify(5)
	a.Modify(5)

	a.String()
}

func sum(a, b int) int {
	return a + b
}

// 多返回值，返回错误码
// 两种错误处理机制
// 1.错误码，类似于java中的检查性异常
// 2.panic，类似于Java中的运行时异常
func sum1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	return a + b, nil
}

// 具名返回值
func sum2(a, b int) (result int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	result = a + b
	err = nil
	return
}

// 可变参数
func sum3(params ...int) int {

	sum := 0
	for _, v := range params {
		sum += v
	}
	return sum
}

// 包级函数
func Sub(a, b int) int {
	return a - b
}

// 闭包
func closure() func() int {
	i := 5
	return func() int {
		// 能访问到函数外部的变量
		// 函数包装数据 => 为函数计算提供了【上下文】
		i++
		return i
	}
}

// 方法
type Age uint

func (a Age) String() {
	fmt.Println("the age is:", a)
}

func (a *Age) Modify(m int) {
	*a = Age(int(*a) + m)
}
