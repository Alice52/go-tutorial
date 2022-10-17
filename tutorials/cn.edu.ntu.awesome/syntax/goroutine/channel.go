package goroutine

import (
	"fmt"
)

func Recv(c chan int) {
	ret, ok := <-c
	fmt.Printf("接收成功: %d, 通道是否关闭: %v\n", ret, ok)
}

func Send(ch chan int) {
	ch <- 10
	fmt.Println("发送成功")
}

func Direction() {
	var ch4 = make(chan int, 1)
	ch4 <- 10
	var ch5 <-chan int // 声明一个只接收通道ch5, 且初始化为 nil
	ch5 = ch4          // 变量赋值时将ch4转为单向通道
	v := <-ch5
	fmt.Printf("v: %v\n", v)
}
