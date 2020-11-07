package doLottery

import (
	"testing"

)

const FILEPATH = "./list.txt"

func TestChooseOne(t *testing.T) {
	//ret:=ChooseOne(FILEPATH)
	ChooseOne(FILEPATH)
	//t.Log(ret)
}
func BenchmarkChooseOne(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		ret :=ChooseOne(FILEPATH)
		b.Log(ret)
	}
}