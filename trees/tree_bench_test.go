package trees_test

import (
	"github.com/norand94/golearn/base"
	"github.com/norand94/golearn/ensemble"
	"testing"
)

func BenchmarkRandomForestFit(b *testing.B) {
	// benchdata.csv contains ../examples/datasets/iris.csv repeated 100 times.
	data, err := base.ParseCSVToInstances("benchdata.csv", true)
	if err != nil {
		b.Fatalf("Cannot load benchdata.csv err:\n%v", err)
	}
	b.ResetTimer()
	tree := ensemble.NewRandomForest(20, 4)
	for i := 0; i < b.N; i++ {
		tree.Fit(data)
	}
}
