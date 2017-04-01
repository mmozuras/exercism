package secret

const testVersion = 1

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(n uint) []string {
	var handshake []string
	for i, action := range actions {
		if n&(uint(1)<<uint(i)) > 0 {
			handshake = append(handshake, action)
		}
	}
	if n&16 > 0 {
		return reverse(handshake)
	}
	return handshake
}

func reverse(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}
