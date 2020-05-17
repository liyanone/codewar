package kata

import "strings"

func SpinWords(str string) string {
	new_str := strings.Fields(str)
	for i, v := range new_str {
		if len(v) >= 5 {
			new_str[i] = reserve(v)
		}
	}
	return strings.Join(new_str, " ")
} // SpinWords

func reserve(str string) string {
	s := make([]byte, len(str))

	for c := 0; c < len(str); c++ {
		s[c] = str[len(str)-1-c]
	}

	return string(s)

}
