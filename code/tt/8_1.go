package tt

import "errors"

func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b cat't be 0")
	}
	return a / b, nil
}

func addsum(nums ...int) int {
	var sum int
	for _, value := range nums {
		sum += value
	}
	return sum
}
