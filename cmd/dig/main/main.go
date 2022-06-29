package main

import (
	"log"

	"github.com/kjushka/di-tests/test_data/action"
)

func main() {
	c1, err := prepareSomeActionWithParams(1, "test", 1.0, 2.0)
	if err != nil {
		log.Fatal(err)
	}
	c2, err := prepareSomeActionWithoutParams(1.0, 2.0)
	if err != nil {
		log.Fatal(err)
	}
	var res1, res2 int
	err = c1.Invoke(func(ss action.SomeService) {
		res1 = ss.MakeJob()
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c2.Invoke(func(ss action.SomeService) {
		res2 = ss.MakeJob()
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result is %v", res1 + res2)
}
