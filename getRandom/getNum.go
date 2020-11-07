package getRandom

import (
	"math/rand"
	"time"
)

func GetNum(max int) int {
	rand.Seed(time.Now().Unix())
	r :=rand.Intn(max)
	time.Sleep(time.Duration(max)*time.Millisecond)
	return r
}
