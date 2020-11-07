package allToSlice

func MakeSlice(m map[string]int)[]string {
	s:=make([]string,0)
	for k,v:=range m{
		for i:=1;i<=v;i++{
			s=append(s,k)
		}
	}
	return s
}
