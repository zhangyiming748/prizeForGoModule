package doLottery

import (
	"github.com/zhangyiming748/prizeForGoModule/getRandom"
	"github.com/zhangyiming748/prizeForGoModule/member"
	"github.com/zhangyiming748/prizeForGoModule/allToSlice"
	"sync"
)

func ChooseOne(s string) string {
	pl:=member.Set2Map(s)

	sl:=allToSlice.MakeSlice(pl)
	var wg sync.WaitGroup
	c1:=make(chan string,1)
	c2:=make(chan string,1)
	c3:=make(chan string,1)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n:=getRandom.GetNum(len(sl))
		v :=sl[n]
		c1<-v
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n:=getRandom.GetNum(len(sl))
		v :=sl[n]
		c2<-v
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n:=getRandom.GetNum(len(sl))
		v :=sl[n]
		c3<-v
	}()
	select {
	case ret := <-c1:
		//fmt.Printf("天选之子是: %s\t天选之子是: %s\t天选之子是: %s\n", ret, ret, ret)
		return ret
	case ret := <-c2:
		//fmt.Printf("天选之子是: %s\t天选之子是: %s\t天选之子是: %s\n", ret, ret, ret)
		return ret
	case ret := <-c3:
		//fmt.Printf("天选之子是: %s\t天选之子是: %s\t天选之子是: %s\n", ret, ret, ret)
		return ret
	}
}