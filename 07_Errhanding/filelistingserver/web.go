package main

import (
	"gogogo/07_Errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

/*知识点串联：
函数式编程的应用：函数作为参数、匿名函数 ，如errWrapper
错误处理和保护：意料之中用error，如文件打不开；意料之外，如数组越界
defer+panic+recover ；try assersion
接口*/

import (
	"gogogo/07_Errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//接口定义即可
type userError interface {
	error
	Message() string
}

//http://localhost:8888/list/n.txt Set Method not allow properties
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError)+"xxx", http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			//此处如果handler返回的err是 userError类型话，会被接管
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//refer to https://zhuanlan.zhihu.com/p/49266481
func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
