package main

import (
	"github.com/kjushka/di-tests/test_data/action"
	"github.com/kjushka/di-tests/test_data/service"

	"github.com/sarulabs/di"
)

func prepareSomeActionWithParams(a int, s string, f1 float32, f2 float64) (di.Container, error) {
	builder, _ := di.NewBuilder()
	err := builder.Add(di.Def{
		Name: "SomeService",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubService1
			c.Fill("SomeSubService1", &ss1)
			var ss2 action.SomeSubService2
			c.Fill("SomeSubService2", &ss2)
			return service.NewSomeServiceWithParam(a, ss1, ss2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubService1",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubSubService1
			c.Fill("SomeSubSubService1", &ss1)
			var ss2 action.SomeSubSubService2
			c.Fill("SomeSubSubService2", &ss2)
			return service.NewSomeSubService1WithParam(s, ss1, ss2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubService2",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubSubService1
			c.Fill("SomeSubSubService1", &ss1)
			return service.NewSomeSubService2WithoutParam(ss1), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubSubService1",
		Build: func(c di.Container) (interface{}, error) {
			return service.NewSomeSubSubService1WithParam(f1), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubSubService2",
		Build: func(c di.Container) (interface{}, error) {
			return service.NewSomeSubSubService2WithParam(f2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	container := builder.Build()
	return container, nil
}

func prepareSomeActionWithoutParams(f1 float32, f2 float64) (di.Container, error) {
	builder, _ := di.NewBuilder()
	err := builder.Add(di.Def{
		Name: "SomeService",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubService1
			c.Fill("SomeSubService1", &ss1)
			var ss2 action.SomeSubService2
			c.Fill("SomeSubService2", &ss2)
			return service.NewSomeServiceWithoutParam(ss1, ss2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubService1",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubSubService1
			c.Fill("SomeSubSubService1", &ss1)
			var ss2 action.SomeSubSubService2
			c.Fill("SomeSubSubService2", &ss2)
			return service.NewSomeSubService1WithoutParam(ss1, ss2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubService2",
		Build: func(c di.Container) (interface{}, error) {
			var ss1 action.SomeSubSubService1
			c.Fill("SomeSubSubService1", &ss1)
			return service.NewSomeSubService2WithoutParam(ss1), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubSubService1",
		Build: func(c di.Container) (interface{}, error) {
			return service.NewSomeSubSubService1WithParam(f1), nil
		},
	})
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name: "SomeSubSubService2",
		Build: func(c di.Container) (interface{}, error) {
			return service.NewSomeSubSubService2WithParam(f2), nil
		},
	})
	if err != nil {
		return nil, err
	}
	container := builder.Build()
	return container, nil
}
