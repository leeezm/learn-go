package main

func main() {
	//12345=>12,345
	println(comma("1234567"))
}

func comma(str string) string {
	length := len(str)
	if length <= 3 {
		return str
	}
	return comma(str[:length-3]) + "," + str[length-3:]
}
