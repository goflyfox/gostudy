package main

import "errors"
import "fmt"

// 按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
func f1(arg int) (int, error) {
	// errors.New 构造一个使用给定的错误信息的基本error 值。
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// 通过实现 Error 方法来自定义 error 类型是可以的。
// 这里使用自定义错误类型来表示上面的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体，并提供了 arg 和 prob 这个两个字段的值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	// 下面的两个循环测试了各个返回错误的函数。
	//  注意在 if行内的错误检查代码，在 Go 中是一个普遍的用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 失败:", e)
		} else {
			fmt.Println("f1 工作:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 失败:", e)
		} else {
			fmt.Println("f2 工作:", r)
		}
	}
	// 你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
