package main

import "fmt"

//import "fmt"

// Retrieve the mount of matching socks colours codes pairs
func RetrieveAmountMatchingSocksPairs(socksAmount uint8, sockColourCodes []uint8) int {

	const limit uint8 = 100

	if socksAmount != uint8(len(sockColourCodes)) {
		return 0
	}

	sockColourCodesLimit := sockColourCodes[0:len(sockColourCodes)]

	if uint8(len(sockColourCodes)) > limit {
		sockColourCodesLimit = sockColourCodes[0 : limit+1]
		socksAmount = limit
	}

	var matches int = 0

	for i := uint8(0); i < socksAmount; i++ {
		if sockColourCodesLimit[i] != 0 {
			for j := i + 1; j < socksAmount; j++ {
				if sockColourCodesLimit[i] == sockColourCodesLimit[j] {
					matches++
					sockColourCodesLimit[j] = 0
					break
				}
			}
		}
	}
	//fmt.Printf("%d", matches)

	return matches
}

func print100() {
	for i := 0; i <= 100; i++ {
		fmt.Printf(",%d", i)
	}
}

func main() {
	print100()
}
