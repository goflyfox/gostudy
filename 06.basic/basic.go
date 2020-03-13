package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// hello world
	/**
	hello world
	*/
	fmt.Println("hello world")

	fmt.Println("##################### values")
	values()

	fmt.Println("##################### variables")
	variables()

	fmt.Println("##################### constants")
	constants()

	fmt.Println("##################### forFunc")
	forFunc()

	fmt.Println("##################### ifElse")
	ifElse()

	fmt.Println("##################### switchFunc")
	switchFunc()
}

// 值
func values() {
	// 字符串拼接用 +
	fmt.Println("hello " + "world " + "!")
	// 整数和浮点数
	fmt.Println("1+2 =", 1+2)
	fmt.Println("11-1 =", 11-1)
	fmt.Println("99*99 =", 99*99)
	fmt.Println("8.0/3.0 =", 8.0/3.0)
	// 布尔型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// 变量
func variables() {
	// var 声明 1 个或者多个变量。
	var a string = "hello"
	fmt.Println(a)
	var b, c int = 3, 5
	fmt.Println(b, c)

	// 会自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)

	// 声明变量且 初始化为0
	var e int
	fmt.Println(e)

	// := 简写会自动推断类型，只能用在初始化
	f := "short"
	fmt.Println(f)
}

// 常量
// 全局常量
const con = "const"

func constants() {
	fmt.Println(con)

	// const 语句可以出现在任何 var 语句可以出现的地方
	const num = 500 * 500 * 500
	// 常数表达式可以执行任意精度的运算
	const num2 = 4e21 / num
	fmt.Println(num2)
	// 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(int64(num2))

	// 当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。
	// 举个例子，这里的 math.Sin函数需要一个 float64 的参数。
	fmt.Println(math.Sin(num))
}

// For循环
func forFunc() {
	// 最常用的方式，带单个循环条件。
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 for 循环。
	for j := 6; j <= 8; j++ {
		fmt.Println(j)
	}

	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。
	for {
		fmt.Println("for...")
		break
	}

	for n := 0; n <= 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// if/else
func ifElse() {
	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// 你可以不要 else 只用 if 语句。
	if 12%4 == 0 {
		fmt.Println("12 is divisible by 4")
	}

	// 在条件语句之前可以有一个语句；
	// 任何在这里声明的变量都可以在所有的条件分支中使用。
	if num := 7; num < 0 {
		fmt.Println(num, "正数")
	} else if num < 10 {
		fmt.Println(num, "小于10")
	} else {
		fmt.Println(num, "其他")
	}

	// 注意，在 Go 中，你可以不适用圆括号，但是花括号是需要的。
	// Go 里没有三目运算符，
	// 所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句。
}

// 分支结构
func switchFunc() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	// 在一个 case 语句中，你可以使用逗号来分隔多个表达式。
	// 在这个例子中，我们很好的使用了可选的default 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("星期天")
	default:
		fmt.Println("工作日")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。
	//  这里展示了 case 表达式是如何使用非常量的。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("12点前")
	default:
		fmt.Println("12点后，包含12点")
	}

	// 这里是一个函数变量
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("什么类型 %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("嘿")
}
