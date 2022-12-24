package day1

type elve struct {
	caloriesOfSupplies []int
	sumOfCalories      int
}

func NewElve() *elve {
	elve := elve{
		caloriesOfSupplies: []int{},
		sumOfCalories:      0,
	}

	return &elve
}
