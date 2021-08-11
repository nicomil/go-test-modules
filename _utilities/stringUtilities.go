package utilities

func StrConcat(s ...string) string {
	str := ""
	for _, elem := range s {
		str += elem + " "
	}
	return str
}
