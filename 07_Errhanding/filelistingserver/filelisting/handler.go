package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//用户具体实现
type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

//return error,use web.go to handler all the error
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		//return error.New("path must start with " + prefix)
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
