package main

import "fmt"

// 我们将通过两个函数：val 和 ptr 来比较指针和值类型的不同。
// val 有一个 int 型参数，所以使用值传递。
// val 将从调用它的那个函数中得到一个 val1 形参的拷贝。
func val(val1 int) {
	val1 = 0
}

// ptr 有一和上面不同的 *int 参数，意味着它用了一个 int指针。
// 函数体内的 *iptr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值。
// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
func ptr(iptr *int) {
	*iptr = 0
}

func main() {
	test := 1
	fmt.Println("initial:", test)
	val(test)
	fmt.Println("val:", test)
	// 通过 &test 语法来取得 test 的内存地址，例如一个变量i 的指针。
	ptr(&test)
	fmt.Println("ptr:", test)
	// 指针也是可以被打印的。
	fmt.Println("pointer:", &test)
	// 	val 在 main 函数中不能改变 test 的值，但是zeroptr 可以，因为它有一个这个变量的内存地址的引用。
	fmt.Println("pointer:", *&test)
}
