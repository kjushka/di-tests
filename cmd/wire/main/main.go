package main

import (
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	res1 := prepareSomeActionWithParams(1, "test", 1.0, 2.0).MakeJob()
	res2 := prepareSomeActionWithoutParams(1.0, 2.0).MakeJob()
	jobTime := time.Since(startTime)
	log.Printf("all wire process ended in %v, results: %v, %v", jobTime, res1, res2)
}
