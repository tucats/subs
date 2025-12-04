package subs

import (
	"math"
	"testing"
)

func Test_getInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value any
		want  int
	}{
		{
			name:  "integer value",
			value: 42,
			want:  42,
		},
		{
			name:  "float value",
			value: 42.0,
			want:  42,
		},
		{
			name:  "float value with fractional part round down",
			value: 42.35,
			want:  42,
		},
		{
			name:  "float value with fractional part round up",
			value: 42.50002,
			want:  43,
		},
		{
			name:  "float value with fractional part that overflows",
			value: 42.0e+70,
			want:  math.MaxInt,
		},
		{
			name:  "string value",
			value: "55",
			want:  55,
		},
		{
			name:  "string value with invalid integer",
			value: "bogus",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getInt(tt.value)
			if got != tt.want {
				t.Errorf("getInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
