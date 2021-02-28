package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var initDone = false

func LogInit(logFolder string) {

	if initDone {
		return
	}
	if logFolder == "" {
		logFolder = "/log"
	}
	var pwd, _ = os.Getwd()
	if _, err := os.Stat(pwd + logFolder); os.IsNotExist(err) {
		err = os.Mkdir(pwd+logFolder, 0777)
		if err != nil {
			fmt.Printf("create log folder: %v\n", err)
			log.Fatal(err)
			os.Exit(-1)
		}
	}
	fmt.Println("[CURRENT PATH] ", pwd)
	pathSep := "/"
	var current_path_list []string = strings.SplitAfter(pwd, pathSep)
	last_index := len(current_path_list) - 1
	if last_index <= 0 {
		log.Fatal("[ERROR] 当前目录不存在测试用例")
		os.Exit(-1)
	}
	goName := current_path_list[last_index]
	log.Printf("caller : %s\n", goName)
	//filePrefix := filepath.Join(pwd+logFolder, goName+"_"+BuildFileName())
	filePrefix := filepath.Join(pwd+logFolder, goName+"_")
	fmt.Println("[CURRENT PATH] ", filePrefix)
	fname := "PickFileName"
	//fname, err := PickFileName(filePrefix)
	//if err != nil {
	//	fmt.Printf("Set log file: %v\n", err)
	//	log.Fatal(err)
	//	os.Exit(-1)
	//}

	f, err1 := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0777)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.SetOutput(f)
	log.Printf("Create log file %v", fname)
	log.SetFlags(log.Ltime | log.Lshortfile)
	initDone = true
}
