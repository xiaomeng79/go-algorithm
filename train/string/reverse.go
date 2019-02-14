package string


//字符串反转(三步反转法)
//将 abcdef 反转成 defabc //1.abc->cba 2.def->fed 3.cbafed->defabc
func LeftRotateStritog(s string,m,n int) string {
	_s := []byte(s)
	reverseString(_s,0,m-1)
	reverseString(_s,m,n-1)
	reverseString(_s,0,n-1)
	return string(_s)
}

func reverseString(s []byte, from,to int) {
	for from < to {
		s[from],s[to] = s[to],s[from]
		from++
		to--
	}
}