package main

import "fmt"

func main() {
	s := make([]byte, 5)
	a := len(s)
	b := cap(s)
	fmt.Printf("len(s)是%d,cap(s)是%d", a, b)
	s = s[2:4]
	a = len(s)
	b = cap(s)
	fmt.Println("")
	fmt.Printf("修改后len(s)是%d,cap(s)是%d", a, b)

	str := "hello，世界"
	c := len(str)
	d := 0
	for i := range str {
		fmt.Println(str[i])
		d += 1
	}
	//len()函数返回的是字节数,不是字符数,每个中文字符是三个字节
	//Go语言中的字符串是UTF-8编码的字节序列，不是字符数组
	//for range按Unicode字符（rune）遍历
	fmt.Printf("该字符串的长度是%d,forrange该字符串的循环次数是%d", c, d)
}
