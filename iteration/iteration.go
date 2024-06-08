package iteration

func Repeat(s string, x int) string {
	var repeated string
	for i := 0; i < x; i++ {
		repeated += s
	}
	return repeated
}
