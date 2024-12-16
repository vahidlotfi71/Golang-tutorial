package units

import "strings"

func ConcatStringWithoutBuilder(a, b string) string {
	var s string = ""
	for i := 0; i < 100; i++ {
		s = s + a + b
	}
	return s
}

func ConcatStringWithBuilder(a, b string) string {
	var sb strings.Builder
	for i := 0; i < 100; i++ {
		sb.WriteString(a+b)
	}
	return sb.String()
}