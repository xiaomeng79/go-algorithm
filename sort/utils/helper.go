package utils

import (
	"os"
	"path/filepath"
	"bufio"
	"strconv"
)

const (
	File_Name = "../utils/intarray.txt"
	Num_Vol = 100000
)

//获取文件绝对路径
func absPath() string {
	_fpath,err := filepath.Abs(File_Name)
	if err != nil{
		panic(err)
	}
	return _fpath
}

//判断随机数文件是否存在
func pathExists() bool {
	_fpath := absPath()
	_,err := os.Stat(_fpath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
//生成随机数文件
func gen() {
	_fpath := absPath()
	GenRandFile(Num_Vol,_fpath)
}

//获取随机数
func GetArrayOfSize(n int) []int {
	_fpath := absPath()
	//判断文件是否存在
	if !pathExists() {
		gen()//不存在创建
	}
	//存在读取
	f ,err := os.Open(_fpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	numbers := make([]int,n)
	scanner := bufio.NewScanner(f)

	for i:=0;i<n;i++ {
		scanner.Scan()
		s,_ := strconv.Atoi(scanner.Text())
		numbers = append(numbers,s)
	}
	return numbers
}

