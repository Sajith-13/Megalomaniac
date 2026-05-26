package combinatorics

import "sort"

func GetCombinationIndex(id1, id2, id3 int) int {
	ids := []int{id1, id2, id3}
	sort.Ints(ids)
	
	x, y, z := ids[0], ids[1], ids[2]

	return choose(z, 3) + choose(y, 2) + choose(x, 1)
}

func choose(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k == 1 {
		return n
	}
	
	if k == 2 {
		return (n * (n - 1)) / 2
	}
	if k == 3 {
		return (n * (n - 1) * (n - 2)) / 6
	}
	return 0
}