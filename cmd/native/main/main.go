package main

import "log"

func main() {
	c1 := prepareSomeActionWithParams(1, "test", 1.0, 2.0)
	c2 := prepareSomeActionWithoutParams(1.0, 2.0)
	var res1, res2 int
	res1 = c1.GetSomeService().MakeJob()
	res2 = c2.GetSomeService().MakeJob()
	log.Printf("result is %v", res1+res2)
}
