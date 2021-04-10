package main

import "fmt"

func main() {
	ch:=make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			ch<-i                     //多个协程并发，需要用Channel操作进行并发的控制
			fmt.Println(i)
		}()
		                               //原代码此处向channel中传入数据，主协程和协程都被阻塞,形成通道死锁
		<-ch
		}
	fmt.Println("over!!!")
}
