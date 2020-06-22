package goderscore

func Clamp(num, lower, upper int) int {
	if num >= lower && num <= upper {
		return num
	}

	lowerDiff := num - lower
	upperDiff := num - upper

	if lowerDiff < 0 {
		lowerDiff = -lowerDiff
	}

	if upperDiff < 0 {
		upperDiff = -upperDiff
	}

	if lowerDiff < upperDiff {
		return lower
	}

	return upper
}

func InRange(num, start, end int) int {

}