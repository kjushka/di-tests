package action

type SomeService interface {
	MakeJob() int
}

type SomeSubService1 interface {
	MakeSubJob1() string
}

type SomeSubService2 interface {
	MakeSubJob2()
}

type SomeSubSubService1 interface {
	MakeSubSubJob1()
}

type SomeSubSubService2 interface {
	MakeSubSubJob2()
}