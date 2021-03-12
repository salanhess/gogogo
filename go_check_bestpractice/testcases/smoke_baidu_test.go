package smoke_test

/*
	用例说明：

*/

import (
	"gogogo/go_check_bestpractice/model"
	"gogogo/go_check_bestpractice/scenario"
	"gogogo/go_check_bestpractice/steps"
	. "gopkg.in/check.v1"
	//"gogogo/go_check_bestpractice/utils"
	"testing"
)

var context *steps.SharedContext = steps.NewSharedContextObject()

//var volName = util.GenRandomName("-_.create")

var params = &model.GetBaiduParams{
	//Hello:       &steps.Envcfg.Greeting,
	Hello: "json",
}

var _ = Suite(&scenario.HelloBaidu{
	Ctx:     context,
	SendReq: params,
	Desc:    "2_smoke_req_test",
})

// 创建云硬盘，并查看云硬盘状态，直到云硬盘状态为available
var _ = Suite(&scenario.HelloBaidu{
	Ctx:     context,
	SendReq: params,
	Desc:    "1_smoke_req_test",
})

// Hook up gocheck into the "go test" runner.
func TestSmokeBaiduScenario(t *testing.T) { TestingT(t) }
