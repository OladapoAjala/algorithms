package dynamicprogramming

func packTile(space int, tiles []int) int {
	memo := make(map[int]int)
	tileFreq := make(map[int]int)
	for _, t := range tiles {
		tileFreq[t] += 1
	}
	return dp(space, tileFreq, memo)
}

func dp(space int, tileFreq, memo map[int]int) int {
	if space < 0 {
		return 0
	}
	if space == 0 {
		return 1
	}

	if cahcedPerm, ok := memo[space]; ok {
		return cahcedPerm
	}

	perm := 0
	for tile, freq := range tileFreq {
		perm += freq * dp(space-tile, tileFreq, memo)
	}
	memo[space] = perm
	return perm
}
