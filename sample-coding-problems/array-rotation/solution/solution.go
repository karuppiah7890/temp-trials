package solution

import "math"

func rotateRight(A []int, K int) []int {
	lastKElementsToPutInFront := A[len(A)-K:]
	return append(lastKElementsToPutInFront, A[:len(A)-K]...)
}

func rotateLeft(A []int, K int) []int {
	firstKElementsToPutInBack := A[:K]
	return append(A[K:], firstKElementsToPutInBack...)
}

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	if K >= len(A) {
		K = len(A) % K
	}

	if K == 0 {
		return A
	}

	KLeft := len(A) - K

	minimumK := math.Min(float64(K), float64(KLeft))

	if minimumK == float64(K) {
		return rotateRight(A, K)
	}

	return rotateLeft(A, KLeft)
}
