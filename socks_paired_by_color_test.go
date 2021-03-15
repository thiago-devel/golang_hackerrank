package main

import "testing"

func TestMatching0Pairs(t *testing.T) {

	const pairs uint8 = 7
	sockColourCodes := []uint8{1, 2, 1, 2, 3, 2}
	var want int = 0

	got := RetrieveAmountMatchingSocksPairs(pairs, sockColourCodes)

	if got != want {
		t.Errorf("got %d, wanted %q", got, want)
	}
}

func TestMatching2Pairs(t *testing.T) {

	const pairs uint8 = 7
	sockColourCodes := []uint8{1, 2, 1, 2, 1, 3, 2}
	var want int = 2

	got := RetrieveAmountMatchingSocksPairs(pairs, sockColourCodes)

	if got != want {
		t.Errorf("got %d, wanted %q", got, want)
	}
}

func TestMatching1Pair(t *testing.T) {

	const pairs uint8 = 4
	sockColourCodes := []uint8{1, 3, 2, 1}
	var want int = 1

	got := RetrieveAmountMatchingSocksPairs(pairs, sockColourCodes)

	if got != want {
		t.Errorf("got %d, wanted %q", got, want)
	}
}

func TestMatching1Upto100Pair(t *testing.T) {

	const socks uint8 = 104
	sockColourCodes := []uint8{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103}
	var want int = 1

	got := RetrieveAmountMatchingSocksPairs(socks, sockColourCodes)

	if got != want {
		t.Errorf("got %d, wanted %q", got, want)
	}
}

func TestMatching1Upto100Pair2(t *testing.T) {

	const socks uint8 = 104
	sockColourCodes := []uint8{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 98, 100, 101, 102, 103}
	var want int = 2

	got := RetrieveAmountMatchingSocksPairs(socks, sockColourCodes)

	if got != want {
		t.Errorf("got %d, wanted %q", got, want)
	}
}
