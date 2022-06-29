package main

import (
	"github.com/kjushka/di-tests/test_data/action"
	"github.com/kjushka/di-tests/test_data/service"
)

type diWithParams struct {
	a  int
	s  string
	f1 float32
	f2 float64
}

func (diwith *diWithParams) GetSomeService() action.SomeService {
	return service.NewSomeServiceWithParam(diwith.a, diwith.GetSomeSubService1(), diwith.GetSomeSubService2())
}

func (diwith *diWithParams) GetSomeSubService1() action.SomeSubService1 {
	return service.NewSomeSubService1WithParam(diwith.s, diwith.GetSomeSubSubService1(), diwith.GetSomeSubSubService2())
}

func (diwith *diWithParams) GetSomeSubService2() action.SomeSubService2 {
	return service.NewSomeSubService2WithoutParam(diwith.GetSomeSubSubService1())
}

func (diwith *diWithParams) GetSomeSubSubService1() action.SomeSubSubService1 {
	return service.NewSomeSubSubService1WithParam(diwith.f1)
}

func (diwith *diWithParams) GetSomeSubSubService2() action.SomeSubSubService2 {
	return service.NewSomeSubSubService2WithParam(diwith.f2)
}

type diWithoutParams struct {
	f1 float32
	f2 float64
}

func (diwithout *diWithoutParams) GetSomeService() action.SomeService {
	return service.NewSomeServiceWithoutParam(diwithout.GetSomeSubService1(), diwithout.GetSomeSubService2())
}

func (diwithout *diWithoutParams) GetSomeSubService1() action.SomeSubService1 {
	return service.NewSomeSubService1WithoutParam(diwithout.GetSomeSubSubService1(), diwithout.GetSomeSubSubService2())
}

func (diwithout *diWithoutParams) GetSomeSubService2() action.SomeSubService2 {
	return service.NewSomeSubService2WithoutParam(diwithout.GetSomeSubSubService1())
}

func (diwithout *diWithoutParams) GetSomeSubSubService1() action.SomeSubSubService1 {
	return service.NewSomeSubSubService1WithParam(diwithout.f1)
}

func (diwithout *diWithoutParams) GetSomeSubSubService2() action.SomeSubSubService2 {
	return service.NewSomeSubSubService2WithParam(diwithout.f2)
}

type DI interface {
	GetSomeService() action.SomeService
	GetSomeSubService1() action.SomeSubService1
	GetSomeSubService2() action.SomeSubService2
	GetSomeSubSubService1() action.SomeSubSubService1
	GetSomeSubSubService2() action.SomeSubSubService2
}

func prepareSomeActionWithParams(a int, s string, f1 float32, f2 float64) DI {
	return &diWithParams{a, s, f1, f2}
}

func prepareSomeActionWithoutParams(f1 float32, f2 float64) DI {
	return &diWithoutParams{f1, f2}
}
