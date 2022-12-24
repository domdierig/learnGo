package day1

import (
	"fmt"

	"github.com/domdierig/learnGo/helper"
)

func CalcCalories() {
	file, error := helper.ReadFileToString("./adventOfCode/day1/input.txt")

	if error != nil {
		fmt.Println(file)
	}
}
