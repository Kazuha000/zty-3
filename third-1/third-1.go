package main

import "fmt"

var ch =make(chan int)
func main()  {
	go func() {
		for i:=0;i<=100;i++ {
			ch<-i
			fmt.Println(i)

		}
	}()

for i:=0;i<=100;i++{
	<-ch
	fmt.Println(i)
}
}