package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var digits []int
	for i := min; i < max; i++ {
		digits = append(digits, i)
	}
	return digits
}
