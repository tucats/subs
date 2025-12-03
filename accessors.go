package subs

import (
	"math"
	"strconv"
)

func getInt(value any) int {
	switch v := value.(type) {
	case int:
		return v

	case float64:
		return int(math.Round(v))

	case string:
		d, err := strconv.Atoi(v)
		if err == nil {
			return d
		}
	}

	return 0
}
