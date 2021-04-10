package main

import (
	"fmt"
	"os"
)
func main() {
	func(){
		f1, _ := os.Create("./proverb.txt")    //创建一个"proverb.txt"文件
		defer f1.Close()
	    proverbs := []string{"don't communicate by sharing memory share memory by communicating\n"}
	    for _, p := range proverbs {
			f1.Write([]byte(p))                 //写入"proverb.txt"中。
		}
	}()

    func(){
    	f2,_:=os.Open("./proverb.txt")
    	defer f2.Close()
    	var tmp [128]byte
    	f2.Read(tmp[:])
    	fmt.Println(string(tmp[:]))       //使用file.Read函数读取"proverb.txt"的内容，并打印到控制台

    }()
}

