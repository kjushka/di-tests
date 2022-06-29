//go:build wireinject
// +build wireinject

package main

import (
	"github.com/kjushka/di-tests/test_data/action"
	"github.com/kjushka/di-tests/test_data/service"

	"github.com/google/wire"
)

func prepareSomeActionWithParams(a int, s string, f1 float32, f2 float64) action.SomeService {
	panic(wire.Build(
		service.NewSomeServiceWithParam,
		service.NewSomeSubService1WithParam,
		service.NewSomeSubService2WithoutParam,
		service.NewSomeSubSubService1WithParam,
		service.NewSomeSubSubService2WithParam,
	))
}

func prepareSomeActionWithoutParams(f1 float32, f2 float64) action.SomeService {
	panic(wire.Build(
		service.NewSomeServiceWithoutParam,
		service.NewSomeSubService1WithoutParam,
		service.NewSomeSubService2WithoutParam,
		service.NewSomeSubSubService1WithParam,
		service.NewSomeSubSubService2WithParam,
	))
}
