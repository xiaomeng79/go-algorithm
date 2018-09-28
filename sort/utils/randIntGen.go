package utils

import (
	"io"
	"math/rand"
	"os"
	"strconv"
)

//生成随机数
func getRandInt(anchor int) int {
	rand.Seed(int64(anchor))
	return rand.Intn(100000)
}

//生成随机数文件
func GenRandFile(num int, filepath string) {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < num; i++ {
		io.WriteString(f, strconv.Itoa(getRandInt(i))+"\n")
	}
}
