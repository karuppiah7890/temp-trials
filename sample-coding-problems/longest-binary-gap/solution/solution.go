package solution

func Solution(N int) int {
	n := N
	noticedOne := false
	zeroCount := 0
	maxBinaryGap := 0

	for n != 0 {
		nextN := n / 2
		remainder := n % 2

		if remainder == 1 {
			if noticedOne {
				binaryGap := zeroCount
				if (binaryGap > maxBinaryGap) {
					maxBinaryGap = binaryGap
				}
				zeroCount = 0
			} else {
				noticedOne = true
			}
		} else {
			if noticedOne {
				zeroCount++
			}
		}

		n = nextN
	}

	return maxBinaryGap
}
