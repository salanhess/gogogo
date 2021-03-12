package scenario

/*
Describe 会先于List执行，与Test写的顺序无关 ，因为case是并行设计
=== RUN   TestSmokeBaiduScenario
2021/03/01 00:09:49 [HelloBaidu][Info][SetUpSuite]1_smoke_req_test
2021/03/01 00:09:49 [HelloBaidu][Info][TestDescribe]1_smoke_req_test {json}
2021/03/01 00:09:49 [HelloBaidu][Info][TestList]1_smoke_req_test
2021/03/01 00:09:49 [HelloBaidu][Info][TearDownSuite]1_smoke_req_test
*/

import (
	"gogogo/go_check_bestpractice/model"
	"gogogo/go_check_bestpractice/steps"
	"log"

	. "gopkg.in/check.v1"
)

type HelloBaidu struct {
	Ctx     *steps.SharedContext
	SendReq *model.GetBaiduParams
	Desc    string
}

func (scenario *HelloBaidu) SetUpSuite(c *C) {
	log.Printf("[HelloBaidu][Info][SetUpSuite]%s\n", scenario.Desc)
}

func (scenario *HelloBaidu) TearDownSuite(c *C) {
	log.Printf("[HelloBaidu][Info][TearDownSuite]%s\n", scenario.Desc)
}

func (scenario *HelloBaidu) TestList(c *C) {
	log.Printf("[HelloBaidu][Info1][TestList]%s\n", scenario.Desc)
}

func (scenario *HelloBaidu) TestDescribe(c *C) {
	log.Printf("[HelloBaidu][Info2][TestDescribe]%s %v \n", scenario.Desc, *scenario.SendReq)

}
