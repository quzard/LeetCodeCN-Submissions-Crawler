func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	res := []byte{}
	var num byte
	for i >= 0 || j >= 0 {
		if i >= 0 {
			num += a[i] - '0'
			i--
		}
		if j >= 0 {
			num += b[j] - '0'
			j--
		}
		res = append(res, num%2+'0')
		num /= 2
	}
	if num > 0 {
		res = append(res, num%2+'0')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}