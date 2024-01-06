package dynamicprogramming

func packTile(space int, tiles []int) float64 {
	// memo := make(map[int]float64)
	return iterative(space, getTileFreq(tiles))
}

func recursive(space int, tileFreq, memo map[int]float64) float64 {
	if space < 0 {
		return 0
	}
	if space == 0 {
		return 1
	}

	if cahcedPerm, ok := memo[space]; ok {
		return cahcedPerm
	}

	var perm float64 = 0
	for tile, freq := range tileFreq {
		perm += freq * recursive(space-tile, tileFreq, memo)
	}
	memo[space] = perm
	return perm
}

func getTileFreq(tiles []int) map[int]float64 {
	tileFreq := make(map[int]float64)
	for _, t := range tiles {
		tileFreq[t] += 1
	}
	return tileFreq
}

func iterative(space int, tileFreq map[int]float64) float64 {
	if space < 0 {
		return 0
	}
	if space == 0 {
		return 1
	}

	dp := make([]float64, space+1)
	dp[0] = 1

	for currSpace := 1; currSpace <= space; currSpace++ {
		var perm float64
		for tile, freq := range tileFreq {
			if tile > currSpace {
				continue
			}
			perm += freq * dp[currSpace-tile]
		}
		dp[currSpace] = perm
	}

	return dp[space]
}
