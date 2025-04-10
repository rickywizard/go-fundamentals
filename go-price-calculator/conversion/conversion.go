package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strs []string) ([]float64, error) {
	prices := make([]float64, len(strs))

	for i, str := range strs {
		floatPrice, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		prices[i] = floatPrice
	}

	return prices, nil
}
