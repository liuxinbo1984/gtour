package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	// 错误处理：可预期的，不太严重的问题
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	s, err := add(-1, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(s)
	// }

	// 错误断言
	if cm, ok := err.(*commonError); ok {
		fmt.Printf("error code:%d, error msg:%s\n", cm.errorCode, cm.errorMsg)
	} else {
		fmt.Println(s)
	}

	// error嵌套： 保留原error信息，防止错误信息丢失
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap一个错误:%w", e)
	fmt.Println(w)

	// 获取原始错误:解包内部错误
	fmt.Println(errors.Unwrap(w))

	// 错误相等判断(a是b)：error和target error相同；error包含target error
	// 替换==类型比较，如： w == e
	fmt.Println("错误相等判断:", errors.Is(w, e))

	// 替换error断言
	s, err1 := add(-4, 2)
	var cm *commonError
	if errors.As(err1, &cm) { //err1信息拷贝给cm(只拷贝相应的类型信息，即使err1是被包装过了)
		fmt.Printf("error code:%d, error msg:%s\n", cm.errorCode, cm.errorMsg)
	} else {
		fmt.Println(s)
	}

	// panic恢复
	// defer的函数会被压入一个defer栈
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	throwPanic()
	fmt.Println("done")
}

// error工厂函数
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		// return 0, errors.New("a或b不能为负数")

		// 自定义error
		// return 0, &commonError{
		// 	errorCode: 1,
		// 	errorMsg:  "a或b不能为负数",
		// }

		// 包装业务错误信息
		return 0, fmt.Errorf("业务异常%w,", &commonError{
			errorCode: 1,
			errorMsg:  "a或b不能为负数",
		})
	}
	return a + b, nil
}

// 自定义error
type commonError struct {
	errorCode int
	errorMsg  string
}

func (c *commonError) Error() string {
	return c.errorMsg
}

func throwPanic() {
	panic("系统错误！")
}
