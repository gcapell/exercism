package secret

var table = map[int][]string{}

func init() {
	for i := 1; i <= 31; i++ {
		table[i] = handshake(i)
	}
}

func handshake(n int) []string {
	var reply []string
	for shift, s := range []string{"wink", "double blink", "close your eyes", "jump"} {
		if n&(1<<uint(shift)) != 0 {
			reply = append(reply, s)
		}
	}
	if n&16 != 0 {
		// reverse
		for j, k := 0, len(reply)-1; j < k; j, k = j+1, k-1 {
			reply[k], reply[j] = reply[j], reply[k]
		}
	}
	return reply
}

func Handshake(n int) []string { return table[n] }
