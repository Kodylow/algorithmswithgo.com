package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}

// func Sum(numbers []int) int {
// 	var s int
// 	for _, i := range numbers {
// 		s += i
// 	}
// 	return s
// }