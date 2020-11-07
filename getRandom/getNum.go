package getRandom

import (
	"math/rand"
	"time"
)

func GetNum1(max int) int {
	rand.Seed(time.Now().Unix())
	r :=rand.Intn(max)
	time.Sleep(time.Duration(max)*time.Millisecond)
	return r
}
func GetNum2(max int)int{
	rand.Seed(time.Now().UnixNano())
	r :=rand.Intn(max)
	time.Sleep(time.Duration(max)*time.Millisecond)
	return r
}
func GetNum3(max int) int {
	rand.Seed(time.Now().Unix()*time.Now().Unix())
	r :=rand.Intn(max)
	time.Sleep(time.Duration(max)*time.Millisecond)
	return r
}