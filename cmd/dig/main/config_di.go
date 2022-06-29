package main

import (
	"github.com/kjushka/di-tests/test_data/action"
	"github.com/kjushka/di-tests/test_data/service"

	"go.uber.org/dig"
)

func prepareSomeActionWithParams(a int, s string, f1 float32, f2 float64) (*dig.Container, error) {
	c := dig.New()
	err := c.Provide(func(ss1 action.SomeSubService1, ss2 action.SomeSubService2) action.SomeService {
		return service.NewSomeServiceWithParam(a, ss1, ss2)
	})
	if err != nil {
		return nil, err
	}
	err = c.Provide(func(ss1 action.SomeSubSubService1, ss2 action.SomeSubSubService2) action.SomeSubService1 {
		return service.NewSomeSubService1WithParam(s, ss1, ss2)
	})
	if err != nil {
		return nil, err
	}
	err = c.Provide(service.NewSomeSubService2WithoutParam)
	if err != nil {
		return nil, err
	}
	err = c.Provide(func() action.SomeSubSubService1 {
		return service.NewSomeSubSubService1WithParam(f1)
	})
	if err != nil {
		return nil, err
	}
	err = c.Provide(func() action.SomeSubSubService2 {
		return service.NewSomeSubSubService2WithParam(f2)
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func prepareSomeActionWithoutParams(f1 float32, f2 float64) (*dig.Container, error) {
	c := dig.New()
	err := c.Provide(service.NewSomeServiceWithoutParam)
	if err != nil {
		return nil, err
	}
	err = c.Provide(service.NewSomeSubService1WithoutParam)
	if err != nil {
		return nil, err
	}
	err = c.Provide(service.NewSomeSubService2WithoutParam)
	if err != nil {
		return nil, err
	}
	err = c.Provide(func() action.SomeSubSubService1 {
		return service.NewSomeSubSubService1WithParam(f1)
	})
	if err != nil {
		return nil, err
	}
	err = c.Provide(func() action.SomeSubSubService2 {
		return service.NewSomeSubSubService2WithParam(f2)
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}
