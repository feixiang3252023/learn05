package main

import "fmt"

func main() {
	var a string = "abc"
	fmt.Println("hello,world", a) //声明一个局部变量，而不去使用它，则会报错
}
