package slice

const testVersion = 1

func All(n int, s string) (result []string) {
	for i := 0; i+n <= len(s); i++ {
		result = append(result, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	ok = n <= len(s)
	if ok {
		first = UnsafeFirst(n, s)
	}
	return
}
