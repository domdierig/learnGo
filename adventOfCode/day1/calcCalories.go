package day1

import (
	"fmt"
	"strings"

	"github.com/domdierig/learnGo/helper"
)

func CalcCalories() {
	file, error := helper.ReadFileToString("./adventOfCode/day1/input.txt")

	if error == nil {
		splittedFile := strings.Split(file, "\n")

		for index, calories := range splittedFile {
			fmt.Println(index, calories)
		}
	}
}
