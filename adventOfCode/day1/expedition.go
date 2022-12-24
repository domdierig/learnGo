package day1

type expidition struct {
	elves                      []elve
	elveWithMostCalories       elve
	elveWithSecondMostCalories elve
	elveWithThridMostCalories  elve
}

func NewExpidition() *expidition {
	expedition := expidition{
		elves:                      []elve{},
		elveWithMostCalories:       *NewElve(),
		elveWithSecondMostCalories: *NewElve(),
		elveWithThridMostCalories:  *NewElve(),
	}

	return &expedition
}
