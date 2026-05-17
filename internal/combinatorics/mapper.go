package combinatorics

import "sort"

// GetCombinationIndex calculates a unique linear index for a triplet of notables.
// It assumes i, j, and k are distinct IDs assigned to your notables pool.
func GetCombinationIndex(i, j, k int) int {
	// 1. Ensure inputs are sorted (i < j < k) for a deterministic hash
	nodes := []int{i, j, k}
	sort.Ints(nodes)
	a, b, c := nodes[0], nodes[1], nodes[2]

	// 2. Combinatorial Number System formula: nCr(a, 1) + nCr(b, 2) + nCr(c, 3)
	return nCr(a, 1) + nCr(b, 2) + nCr(c, 3)
}

// nCr calculates the binomial coefficient (n choose r)
func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r == 0 || r == n {
		return 1
	}
	if r > n/2 {
		r = n - r
	}
	res := 1
	for i := 1; i <= r; i++ {
		res = res * (n - r + i) / i
	}
	return res
}