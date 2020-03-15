package main

import "fmt"

// 当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。
// 这个特性提升了程序的类型安全性。
func ping(pings chan<- string, msg string) {
	// ping 函数定义了一个只允许发送数据的通道。
	// 尝试使用这个通道来接收数据将会得到一个编译时错误。
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
