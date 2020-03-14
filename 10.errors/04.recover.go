package main

import "fmt"

func main() {
	// 这里我们对异常进行了捕获
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// 我们将在这个网站中使用 panic 来检查预期外的错误。这个是唯一一个为 panic 准备的例子。
	panic("一个异常")

}
