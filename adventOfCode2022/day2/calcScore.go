package day2

import (
	"fmt"
	"strings"

	"github.com/domdierig/learnGo/helper"
)

func CalcScore() {
	file, error := helper.ReadFileToString("./adventOfCode2022/day2/input.txt")

	totalScorePartA := 0
	totalScorePartB := 0

	scoreBoardA := getScoreBoard([3]int{1, 2, 3}, [3][3]int{{3, 6, 0}, {0, 3, 6}, {6, 0, 3}})
	scoreBoardB := getScoreBoard([3]int{0, 3, 6}, [3][3]int{{3, 1, 2}, {1, 2, 3}, {2, 3, 1}})

	if error == nil {
		fileSplitted := strings.Split(file, "\n")
		for _, line := range fileSplitted {
			opponentSign := string(line[0])
			meSign := string(line[2])

			opponentIndex := translateToIndex(opponentSign)
			meIndex := translateToIndex(meSign)

			totalScorePartA += scoreBoardA[opponentIndex][meIndex]
			totalScorePartB += scoreBoardB[opponentIndex][meIndex]
		}
	}

	fmt.Println(scoreBoardA)
	fmt.Println(scoreBoardB)

	fmt.Println(totalScorePartA)
	fmt.Println(totalScorePartB)
}

func getScoreBoard(array [3]int, matrix [3][3]int) [3][3]int {
	result := [3][3]int{{}, {}, {}}

	for i := 0; i < 3; i++ {
		for m := 0; m < 3; m++ {
			result[m][i] = array[i] + matrix[m][i]
		}
	}

	return result
}

func translateToIndex(sign string) int {
	if sign == "A" || sign == "X" {
		return 0
	} else if sign == "B" || sign == "Y" {
		return 1
	}
	return 2
}
