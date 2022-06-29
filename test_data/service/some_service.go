package service

import (
	"math/rand"
	"time"

	"github.com/kjushka/di-tests/test_data/action"
)

type someServiceImpl struct {
	a    int
	sub1 action.SomeSubService1
	sub2 action.SomeSubService2
}

func NewSomeServiceWithParam(
	a int,
	sub1 action.SomeSubService1,
	sub2 action.SomeSubService2,
) action.SomeService {
	return &someServiceImpl{a, sub1, sub2}
}

func NewSomeServiceWithoutParam(sub1 action.SomeSubService1, sub2 action.SomeSubService2) action.SomeService {
	return &someServiceImpl{rand.Int(), sub1, sub2}
}

func (ssp *someServiceImpl) MakeJob() int {
	startTime := time.Now()
	//log.Printf("some service job starts\na value is '%v'", ssp.a)
	res := ssp.sub1.MakeSubJob1()
	ssp.sub2.MakeSubJob2()
	jobTime := time.Since(startTime)
	//log.Printf("some service job finished, result time is %v", jobTime)
	return ssp.a + len(res) + int(jobTime)
}
