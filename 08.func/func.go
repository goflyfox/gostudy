package main

import "fmt"

func main() {
	// hello world
	fmt.Println("hello world")

	// 1. 加法
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// 2. 多值返回
	// 这里我们通过多赋值 操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// 如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
	_, c := vals()
	fmt.Println(c)

	// 3. 可变参数
	// 变参函数使用常规的调用方式，除了参数比较特殊。
	sum(1, 2)
	sum(1, 2, 3)
	// 如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// 4. 闭包
	// 我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。
	// 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
	nextInt := intSeq()
	// 通过多次调用 nextInt 来看看闭包的效果。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
	newInts := intSeq()
	fmt.Println(newInts())

	// 5. 递归
	fmt.Println(fact(7))
}

// 函数
// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
	return a + b
}

// (int, int) 在这个函数中标志着这个函数返回 2 个 int。
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 多返回值函数
func vals() (int, int) {
	return 3, 7
}

// 变参函数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包
// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。
// 这个返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 递归
// face 函数在到达 face(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
