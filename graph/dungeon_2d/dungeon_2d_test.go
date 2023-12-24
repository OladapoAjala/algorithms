package dungeon2d

import "testing"

func Test_dungeon(t *testing.T) {
	input := [][]string{
		{".", "s", ".", "#", ".", ".", "."},
		{"#", "#", ".", ".", ".", "#", "."},
		{"e", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", "#", ".", ".", "."},
		{"#", ".", "#", ".", ".", "#", "."},
		// {".", ".", ".", "#", ".", ".", "."},
		// {"#", ".", ".", "#", ".", ".", "."},
	}
	dungeon(input, 0, 1)
}
