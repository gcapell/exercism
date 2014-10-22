package slice

func All(n int, s string) []string {
	var reply []string

	for j := 0; j+n <= len(s); j++ {
		reply = append(reply, s[j:j+n])
	}
	return reply
}

func Frist(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
