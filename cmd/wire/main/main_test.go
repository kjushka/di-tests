package main

import "testing"

func BenchmarkWireConfigAndJob(b *testing.B) {
	b.ReportAllocs()
	var res1, res2 int
	for i := 0; i < b.N; i++ {
		res1 = prepareSomeActionWithParams(1, "test", 1.0, 2.0).MakeJob()
		res2 = prepareSomeActionWithoutParams(1.0, 2.0).MakeJob()
		if res1 + res2 == 0 {
			b.Log("results sum is equal 0")
		}
	}
}