package doLottery

import (
	"fmt"
	"log"
	"sync"

	"github.com/zhangyiming748/prizeForGoModule/allToSlice"
	"github.com/zhangyiming748/prizeForGoModule/getRandom"
	"github.com/zhangyiming748/prizeForGoModule/member"
)

func ChooseOne(s string) string {
	pl := member.Set2Map(s)

	sl := allToSlice.MakeSlice(pl)
	var wg sync.WaitGroup
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n := getRandom.GetNum1(len(sl))
		log.Println("选中的随机数是: ",n)
		v := sl[n]
		c1 <- v
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n := getRandom.GetNum2(len(sl))
		log.Println("选中的随机数是: ",n)
		v := sl[n]
		c2 <- v
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		n := getRandom.GetNum3(len(sl))
		log.Println("选中的随机数是: ",n)
		v := sl[n]
		c3 <- v
	}()
	select {
	case ret := <-c1:
		fmt.Printf("中奖的这个逼是: %s,大家一起来鄙视他\n", ret)
		return ret
	case ret := <-c2:
		fmt.Printf("中奖的这个逼是: %s,大家一起来鄙视他\n", ret)
		return ret
	case ret := <-c3:
		fmt.Printf("中奖的这个逼是: %s,大家一起来鄙视他\n", ret)
		return ret
	}
}
