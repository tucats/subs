package subs

import (
	"math"
	"strings"
)

// Handle the special case where the format is a decimal/integer format, and the
// value is a floating value that could be a precise decimal. This is a side-effect
// of JSON unmarshaling which assigns a float64 to all numeric fields.
func normalizeForFormat(format string, value any) any {
	format = strings.TrimSpace(format)

	if strings.HasPrefix(format, "%") && strings.HasSuffix(format, "d") {
		if f, ok := value.(float64); ok {
			if math.Round(f) == f && math.Abs(f) < float64(math.MaxInt-1) {
				return int(f)
			}
		}
	}

	return value
}

// isZeroValue returns true if the arbitrary value is considered the "zero state" for
// the given type. That is, if it is a numeric value equal to zero, a boolean value
// equal to false, and empty string, or a map or array with no elements.
func isZeroValue(value any) bool {
	switch v := value.(type) {
	case []any:
		if len(v) == 0 {
			return true
		}

	case []int:
		if len(v) == 0 {
			return true
		}

	case []string:
		if len(v) == 0 {
			return true
		}

	case map[string]any:
		if len(v) == 0 {
			return true
		}

	case map[string]string:
		if len(v) == 0 {
			return true
		}

	case string:
		if v == "" {
			return true
		}

	case int, int32, int64, float32, float64:
		if v == 0 {
			return true
		}

	case bool:
		if !v {
			return true
		}

	case nil:
		return true
	}

	return false
}

// normalizeNumericValues converts numeric values to be either int or float64 values, based
// on the "wantFloat" flag. This is used to convert JSON-marshalled values (usually float64)
// to expected numeric types for formatting by the substitution processor.
//
// If the numeric value would overflow as an int value and the wantFlat flag is false, the
// value is returned as math.MaxInt.
func normalizeNumericValues(value any, wantFloat bool) any {
	switch v := value.(type) {
	case int:
		if wantFloat {
			return float64(v)
		}

		return v

	case int32:
		if wantFloat {
			return float64(v)
		}

		if int64(v) > math.MaxInt {
			return math.MaxInt
		}

		return int(v)

	case int64:
		if wantFloat {
			return float64(v)
		}

		if v > math.MaxInt {
			return math.MaxInt
		}

		return int(v)

	case float32:
		return normalizeNumericValues(float64(v), wantFloat)

	case float64:
		if wantFloat {
			return float64(v)
		}

		vv := math.Abs(v)
		if vv == math.Floor(vv) && v <= math.MaxInt {
			return int(v)
		}

		return v

	default:
		return value
	}
}
