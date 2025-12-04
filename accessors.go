package subs

import (
	"math"
	"strconv"
)

// getInt attempts to convert an arbitrary value into an integer. This is used, for example, to
// get numeric values expressed in a string substitution operator. If the value is a float64
// value and is too large for an int data type, MaxInt is returned. If the value is a string and
// cannot be parsed as an integer, 0 is returned.
func getInt(value any) int {
	switch v := value.(type) {
	case int:
		return v

	case float64:
		vv := math.Round(v)
		if vv > math.MaxInt {
			return math.MaxInt
		}

		return int(vv)

	case string:
		d, err := strconv.Atoi(v)
		if err == nil {
			return d
		}
	}

	return 0
}
