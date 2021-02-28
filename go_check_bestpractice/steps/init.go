package steps

import (
	"flag"
	"fmt"
	"os"
)

//var EnvCfg model.EnvConfig
var TEST_ENV_TYPE string
var pwd, _ = os.Getwd()

//var pathSep = filepath.FromSlash(string(filepath.Separator))
var test_env = os.Getenv("TEST_ENV_CONF")

//var data_file = os.Getenv("DATA_FILE")

var Gopath = os.Getenv("GOPATH")
var EnvCfgPath string = Gopath + "/src/jd.com/zbs/zbs-test/disk_auto/env_settings/" + test_env

//var DataRepoPath string = Gopath + "/src/jd.com/zbs/zbs-test/disk_auto/data_settings/" + data_file

func init() {
	fmt.Println("========================= [" + pwd + "] =========================")
	if os.Getenv("OPENAPI_ENV") != "" && os.Getenv("TEST_ENV") == "" {
		//EnvCfg = OpenapiCfg.PreEnv
		//DataRepo = OpenapiCfg.PreData
	}
	flag.Parse()
	//fmt.Printf("[After parse flag] Dir %s  File %s\n", *Dir, *CaseName)
	env_init()
	//data_init()
}

func env_init() {
	var EnvName string
	if os.Getenv("OPENAPI_ENV") == "" {
		EnvName = os.Getenv("TEST_ENV")
		if "" == EnvName {
			EnvName = "test_env_sq_zbs2"
		}
		info := fmt.Sprintf("=====[init env]=====Use config %s file %s\n", EnvName, EnvCfgPath)
		fmt.Printf(info)
		//err := EnvCfg.GetEnvInfo(EnvName, EnvCfgPath)
		//if err != nil {
		//	fmt.Printf("TEST_ENV %s\n", EnvCfgPath)
		//	log.Fatal(err)
		//}
	}
	//log_init()
	//util.LogInit(EnvCfg.Log)
	//Step = time.Duration(EnvCfg.Step) * time.Millisecond
	//get test_env_type
	TEST_ENV_TYPE = os.Getenv("TEST_ENV_TYPE")
}
