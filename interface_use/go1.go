package main

import (
	"fmt"
	"gogogo/interface_use/mock"
	"gogogo/interface_use/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

//方式一：使用switch 找到对应接口的实际实现类型
func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contend:", v.Contents)
	case *real.Retriever:
		fmt.Println("* Contend:", v.TimeOut)
	}
}
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is a fack imooc.com"}
	fmt.Println(download(r))

	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute + time.Second,
	}
	//方式二，使用type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.UserAgent)
	//fmt.Println(download(r))
	inspect(r)

	//使用类型断言 加上判断
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever!")
	}
}
