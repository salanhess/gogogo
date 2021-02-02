package main

import (
	"bufio"
	"fmt"
	"os"
)

//err handle  1 2 3 to detail
func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("New err")
	if err != nil {
		fmt.Println("err :", err.Error())
		fmt.Println("[err2] :", err)
		if PathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println("[Err3]", PathError.Op, PathError.Path, PathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, "I'm"+string(i))
	}

}
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occured")
	return
	fmt.Println(4)
}

func main() {
	tryDefer()
	writeFile("abc.txt")
}
