package main

import (
	"testing"
)

func BenchmarkNativeConfigAndJob(b *testing.B) {
	b.ReportAllocs()
	var res1, res2 int
	for i := 0; i < b.N; i++ {
		c1 := prepareSomeActionWithParams(1, "test", 1.0, 2.0)
		c2 := prepareSomeActionWithoutParams(1.0, 2.0)
		res1 = c1.GetSomeService().MakeJob()
		res2 = c2.GetSomeService().MakeJob()
		if res1+res2 == 0 {
			b.Log("results sum is equal 0")
		}
	}
}
