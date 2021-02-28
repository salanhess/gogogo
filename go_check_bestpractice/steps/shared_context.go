package steps

import "gogogo/go_check_bestpractice/model"

type SharedContext struct {
	FeedbackMap map[string]*model.GetBaiduResponse
	SkipTest    bool
}

func NewSharedContextObject() *SharedContext {
	s := &SharedContext{}
	s.FeedbackMap = make(map[string]*model.GetBaiduResponse)
	return s
}
