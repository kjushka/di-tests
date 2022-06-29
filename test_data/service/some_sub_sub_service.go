package service

import (
	"time"

	"github.com/kjushka/di-tests/test_data/action"
)

type someSubSubService1Impl struct {
	f float32
}

func NewSomeSubSubService1WithParam(f float32) action.SomeSubSubService1 {
	return &someSubSubService1Impl{f}
}

func (ssss *someSubSubService1Impl) MakeSubSubJob1() {
	// log.Println("some sub sub service 1 job starts")
	// log.Printf("some sub service 1 job finished")
	time.Sleep(time.Microsecond * 50)
}

type someSubSubService2Impl struct {
	f float64
}

func NewSomeSubSubService2WithParam(f float64) action.SomeSubSubService2 {
	return &someSubSubService2Impl{f}
}

func (ssss *someSubSubService2Impl) MakeSubSubJob2() {
	//log.Println("some sub sub service 2 job starts")
	//log.Printf("some sub service 2 job finished")
	time.Sleep(time.Microsecond * 50)
}
