package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	rev := make([]byte, len(word))

	for i := 0; i < len(word); i++ {
		rev[len(rev)-i-1] = word[i]
	}

	return string(rev)
}
