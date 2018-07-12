package shell

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/xiaomeng79/go-algorithm/sort/testdata"
	"github.com/xiaomeng79/go-algorithm/sort/utils"
)

func TestShell(t *testing.T) {
	for _, v := range testdata.Values {
		assert.Exactly(t, v.Sort, ShellSort(v.Nosort), "no eq")
	}
}


func benchmarkShellSort(n int,b *testing.B) {
	b.StopTimer()
	list := utils.GetArrayOfSize(n)
	b.ReportAllocs()
	b.StartTimer()
	for i:=0;i < b.N;i++ {
		ShellSort(list)
	}
}

func BenchmarkShellSort100(b *testing.B) { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }
