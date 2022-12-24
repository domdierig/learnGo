package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/domdierig/learnGo/helper"
)

func CalcCalories() {
	file, error := helper.ReadFileToString("./adventOfCode/day1/input.txt")

	if error == nil {
		splittedFile := strings.Split(file, "\n")
		expidition := *NewExpidition()
		newElve := NewElve()

		for _, caloriesOfSupply := range splittedFile {
			if caloriesOfSupply == "" {
				expidition.elves = append(expidition.elves, *newElve)

				if newElve.sumOfCalories > expidition.elveWithMostCalories.sumOfCalories {
					expidition.elveWithThridMostCalories = expidition.elveWithSecondMostCalories
					expidition.elveWithSecondMostCalories = expidition.elveWithMostCalories
					expidition.elveWithMostCalories = *newElve
				} else if newElve.sumOfCalories > expidition.elveWithSecondMostCalories.sumOfCalories {
					expidition.elveWithThridMostCalories = expidition.elveWithSecondMostCalories
					expidition.elveWithSecondMostCalories = *newElve
				} else if newElve.sumOfCalories > expidition.elveWithThridMostCalories.sumOfCalories {
					expidition.elveWithThridMostCalories = *newElve
				}

				newElve = NewElve()
			} else {
				caloriesOfSupplyInt, error := strconv.Atoi(caloriesOfSupply)

				if error == nil {
					newElve.caloriesOfSupplies = append(newElve.caloriesOfSupplies, caloriesOfSupplyInt)
					newElve.sumOfCalories += caloriesOfSupplyInt
				}
			}
		}

		fmt.Println(expidition.elveWithMostCalories.sumOfCalories)
		fmt.Println(expidition.elveWithMostCalories.sumOfCalories +
			expidition.elveWithSecondMostCalories.sumOfCalories +
			expidition.elveWithThridMostCalories.sumOfCalories)
	}
}
