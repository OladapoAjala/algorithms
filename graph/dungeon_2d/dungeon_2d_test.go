package dungeon2d

import "testing"

func Test_dungeon(t *testing.T) {
	input := [][]string{
		{".", ".", ".", "#", ".", ".", "."},
		{"#", "#", ".", ".", ".", "#", "."},
		{"e", "#", ".", "s", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "."},
		{"#", ".", "#", ".", ".", "#", "."},
		// {".", ".", ".", "#", ".", ".", "."},
		// {"#", ".", ".", "#", ".", ".", "."},
	}
	dungeon(input, 2, 3)
}
