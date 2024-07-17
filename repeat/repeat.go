package repeat

// Repeat takes a character and the times that the character should be repeated and returns a string
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
