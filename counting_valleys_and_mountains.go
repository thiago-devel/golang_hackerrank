package main

import (
	"strings"
)

func CountValleysInPath(steps uint32, path string) int32 {
	var seaLevel int = 0
	var valleys int32 = 0

	for _, step := range path {

		sUpperStep := strings.ToUpper(string(step))

		if sUpperStep == "U" {
			seaLevel = seaLevel + 1
		} else if sUpperStep == "D" {
			seaLevel = seaLevel - 1
		} else { // path contains invalid steps, aborting a counting
			return int32(0)
		}

		if (sUpperStep == "U") && (seaLevel == 0) {
			valleys++
		}
	}

	return valleys
}
