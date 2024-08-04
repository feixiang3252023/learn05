package main

import "fmt"

//iota的用法
// (/* */ 多行注释，//单行注释)
/*func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值, iota += 1
		e        //"ha"    iota += 1
		f = 100  //iota += 1
		g        //100 iotq += 1
		h = iota //7 ,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}*/

// iota的一个有趣的实例
const (
	i = 1 << iota // (<<是左移的意思)
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
