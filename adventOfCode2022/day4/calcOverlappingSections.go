package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/domdierig/learnGo/helper"
)

func CalcOverlappingSections() {
	fmt.Println("Overlapping")

	file, error := helper.ReadFileToString("./adventOfCode2022/day4/input.txt")

	if error == nil {
		fileSplitted := strings.Split(file, "\n")

		sumTotalOverlapping := 0
		sumPartlyOverlapping := 0

		for _, line := range fileSplitted {
			lineSplitted := strings.Split(line, ",")

			front := lineSplitted[0]
			back := lineSplitted[1]

			frontSplitted := strings.Split(front, "-")
			backSplitted := strings.Split(back, "-")

			leftElveLower := convertToInt(frontSplitted[0])
			leftElveUpper := convertToInt(frontSplitted[1])

			rightElveLower := convertToInt(backSplitted[0])
			rightElveUpper := convertToInt(backSplitted[1])

			if (leftElveLower <= rightElveLower && leftElveUpper >= rightElveUpper) || (rightElveLower <= leftElveLower && rightElveUpper >= leftElveUpper) {
				sumTotalOverlapping++
			}

		}

		fmt.Println(sumTotalOverlapping)
		fmt.Println(sumPartlyOverlapping)
	}
}

func convertToInt(s string) int {
	result, error := strconv.Atoi(s)

	if error == nil {
		return result
	}

	return 0
}
