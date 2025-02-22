package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var res string

	for _, r := range word {
		res = string(r) + res
	}
	return res
}

	// var res string

	// for i := len(word) - 1; i >= 0; i-- {
	// 	res += string(word[i])
	// }

	// return res