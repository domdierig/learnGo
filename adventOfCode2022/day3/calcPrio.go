package day3

import (
	"fmt"
	"strings"

	"github.com/domdierig/learnGo/helper"
)

func CalcPrio() {
	file, error := helper.ReadFileToString("./adventOfCode2022/day3/input.txt")

	prioSum := 0

	if error == nil {
		fileSplitted := strings.Split(file, "\n")

		for _, line := range fileSplitted {
			lineLength := len(line)

			front := line[0 : (lineLength/2)-1]
			back := line[(lineLength / 2):(lineLength - 1)]

			for _, character := range front {
				if strings.Contains(back, string(character)) {
					prio := getPrioOfCharacter(character)

					prioSum += prio

					break
				}
			}
		}
	}

	fmt.Println(prioSum)
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.

// a 97
// A 65

func getPrioOfCharacter(character rune) int {
	characterStr := string(character)
	characterStrUpper := strings.ToUpper(characterStr)

	if characterStr == characterStrUpper {
		fmt.Println("Character", string(character), "ASCII Code", character, "Prio Value", int(character)-38)
		return int(character) - 38
	} else {
		fmt.Println("Character", string(character), "ASCII Code", character, "Prio Value", int(character)-96)
		return int(character) - 96
	}
}
