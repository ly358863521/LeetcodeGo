func convertToTitle(n int) string {
	res := ""
	var y int
	for n > 0 {
		n--
		y = n % 26
		n = n / 26
		res = string(byte(65+y)) + res
	}
	return res
}

func convertToTitle(n int) string {
	var res []rune
	var y int
	for n > 0 {
		n--
		y = n % 26
		n = n / 26
		res = append([]rune{rune(65 + y)}, res...)
	}
	return string(res)
}