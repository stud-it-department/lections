package sort

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

func swap(data []int, first, second int) {
	f := data[first]

	data[first] = data[second]
	data[second] = f
}

func Sort(data []int) error {
	if len(data) == 0 {
		return fmt.Errorf("invalid data: %w", ErrInvalidArgument)
	}

	for i, _ := range data {
		minI := i

		for j := i; j < len(data); j++ {
			if data[j] < data[minI] {
				minI = j
			}
		}

		swap(data, i, minI)
	}

	return nil
}
