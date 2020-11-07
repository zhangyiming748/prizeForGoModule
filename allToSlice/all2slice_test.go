package allToSlice

import "testing"

func TestMakeSlice(t *testing.T) {
	m := make(map[string]int)
	m["张三"] = 3
	m["cvcv"] = 1
	s:=MakeSlice(m)
	t.Logf("%s",s)
}
