package main

import (
	"fmt"
	"gogogo/common"
)

/*
裹不清 byte ,uint8 ,string [] 相互的转换
https://zhuanlan.zhihu.com/p/270626496
1.可以使用标准转换的原因：
可以看到，入参str指针就是指向byte的指针，那么我们可以确定string的底层数据结构就是byte数组。

2. https://studygolang.com/articles/15455
GO源码默认用Unicode字符集， 使用UTF-8编码

【存储单位 字节】
计算机存储信息的最小单位，称之为位 bit，二进制的一个0或1叫一位
计算机存储容量基本单位是字节 Byte，8个二进制位组成 1 个字节

【信息表示单位 字符】
字符 是一种符号，像 英文a和中文阿 就是不同字符
不同的字符在不同的编码格式下，所需要的存储单位不一样
ASCLII 编码中一个英文字母一字节，一个汉字两字节
UTF-8 编码中 一个英文字母一字节，一个常见汉字3字节，不常用的超大字符集汉字4字节

Go 源码文件默认采用Unicode字符集，Unicode码点和内存中字节序列的变换实现使用了UTF-8，
这使得Go编程无需考虑编码转换的问题非常方便

从编码上来分析
byte用来强调一个字节代表的数据（例如字符 a 就是 97），而不是数字；
rune用来表示Unicode的码点，即一个字符

通俗一点
byte 只能操作简单的字符，不支持中文操作
rune 能操作任何字符
*/

func conv() {
	//标准转换模式
	c := []byte("hello")
	common.PrintVal(c)
	d := string(c)
	common.PrintVal(d)

	var a []byte
	a = []byte{'a', 'b'}
	common.PrintVal(a)

	a_real := []int8{'a', 'b'}
	common.PrintVal(a_real)

	b := []byte{111, 22}
	common.PrintVal(b)

	e := []rune{11, 22}
	common.PrintVal(e)

	f := []uint32{11, 22}
	common.PrintVal(f)

	g := []rune("你好")
	common.PrintVal(g)

	h := []byte("你好")
	common.PrintVal(h)
}
func main() {
	//conv()
	str := "hello 世界！"

	fmt.Println(str) //会输出 hello 世界!，这证明 Go 是UTF-8 编码的，输出长度为 13 这说明了一个汉字3字节
	fmt.Println(len(str))
	fmt.Println(str[1])
	fmt.Println(string(str[1]))
	fmt.Println(str[1:]) //输出 ello 世界! 说明 string 底层的数据结构是数组
	fmt.Println(str[7:]) //输出 ��界! 说明 string 底层是一个byte 数组，不然不会乱码

	x := []rune(str)
	fmt.Println(len(x))
	common.PrintVal((x[6:]))
	common.PrintVal(string(x[6:]))
}
