package main

import (
	"testing"

	"github.com/kjushka/di-tests/test_data/action"
)

func BenchmarkDigConfigAndJob(b *testing.B) {
	b.ReportAllocs()
	var res1, res2 int
	for i := 0; i < b.N; i++ {
		c1, err := prepareSomeActionWithParams(1, "test", 1.0, 2.0)
		if err != nil {
			b.Fatal(err)
		}
		c2, err := prepareSomeActionWithoutParams(1.0, 2.0)
		if err != nil {
			b.Fatal(err)
		}
		err = c1.Invoke(func(ss action.SomeService) {
			res1 = ss.MakeJob()
		})
		if err != nil {
			b.Fatal(err)
		}
		err = c2.Invoke(func(ss action.SomeService) {
			res2 = ss.MakeJob()
		})
		if err != nil {
			b.Fatal(err)
		}
		if res1 + res2 == 0 {
			b.Log("results sum is equal 0")
		}
	}
}
