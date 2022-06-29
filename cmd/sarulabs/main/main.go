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
	res1 = c1.Get("SomeService").(action.SomeService).MakeJob()
	res2 = c2.Get("SomeService").(action.SomeService).MakeJob()
	log.Printf("result is %v", res1+res2)
}
