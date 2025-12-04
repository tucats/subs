package subs

import (
	"fmt"
	"sort"
)

// Convert an arbitrary object into an array of strings, using the supplied
// Go format string as needed.
//
// A map becomes and array of key: value pairs expressed as a string value.
// Arrays become an array of string values.
func makeArray(values any, format string) []string {
	var result []string

	switch v := values.(type) {
	case map[string]any:
		format = "%s: " + format
		keys := make([]string, 0, len(v))

		for key := range v {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			result = append(result, fmt.Sprintf(format, key, v[key]))
		}

	case map[string]string:
		format = "%s: " + format
		keys := make([]string, 0, len(v))

		for key := range v {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			result = append(result, fmt.Sprintf(format, key, v[key]))
		}

	case []any:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []int:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []int32:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []int64:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []float32:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []float64:
		for _, item := range v {
			result = append(result, fmt.Sprintf(format, item))
		}

	case []string:
		result = v

	// It wasn't an array, so just make a single-element array
	// with the value in it as a string.
	default:
		result = []string{fmt.Sprintf("%v", v)}
	}

	return result
}
