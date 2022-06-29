package service

import (
	"fmt"

	"github.com/kjushka/di-tests/test_data/action"
)

type someSubService1Impl struct {
	s    string
	sub1 action.SomeSubSubService1
	sub2 action.SomeSubSubService2
}

func NewSomeSubService1WithParam(
	s string,
	sub1 action.SomeSubSubService1,
	sub2 action.SomeSubSubService2,
) action.SomeSubService1 {
	return &someSubService1Impl{s, sub1, sub2}
}

func NewSomeSubService1WithoutParam(
	sub1 action.SomeSubSubService1,
	sub2 action.SomeSubSubService2,
) action.SomeSubService1 {
	return &someSubService1Impl{"Hello World", sub1, sub2}
}

func (sss *someSubService1Impl) MakeSubJob1() string {
	//log.Printf("some sub service 1 job starts\ns value is '%s'", sss.s)
	sss.sub1.MakeSubSubJob1()
	sss.sub2.MakeSubSubJob2()
	//log.Printf("some sub service 1 job finished")
	return fmt.Sprintf("%s %v", sss.s, len(sss.s))
}

type someSubService2Impl struct {
	sub1 action.SomeSubSubService1
}

func NewSomeSubService2WithoutParam(
	sub1 action.SomeSubSubService1,
) action.SomeSubService2 {
	return &someSubService2Impl{sub1}
}

func (sss *someSubService2Impl) MakeSubJob2() {
	//log.Println("some sub service 2 job starts")
	sss.sub1.MakeSubSubJob1()
	//log.Printf("some sub service 2 job finished")
}
