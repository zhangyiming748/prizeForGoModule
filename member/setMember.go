package member

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Set2Map(s string) map[string]int {
	m := make(map[string]int)
	f, err := os.Open(s)
	if err != nil {
		fmt.Printf("读取文件发生错误,报错信息是: %v", err)
	}
	defer func() {
		_ = f.Close()
	}()
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		//a=[]uint8
		if c == io.EOF {
			//log.Println("录入结束")
			break
		}
		result := uint8ToString(a)
		if _, ok := m[result]; ok {
			m[result] = m[result] + 1
		} else {
			m[result] = 1
		}
	}
	return m
}
func uint8ToString(u8 []uint8) string {
	strout := string(u8)
	return strout
}
