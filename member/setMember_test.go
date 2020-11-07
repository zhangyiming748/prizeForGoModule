package member

import (
	"testing"
)

func TestSet2Map(t *testing.T) {
	path:="./list.txt"
	m:=Set2Map(path)
	t.Log(m)
}
func BenchmarkSet2Map(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		path:="./list.txt"
		m:=Set2Map(path)
		b.Log(m)
	}
	b.StopTimer()
}