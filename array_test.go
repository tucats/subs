package subs

import (
	"reflect"
	"testing"
)

func Test_makeArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values any
		format string
		want   []string
	}{
		{
			name:   "Simple array of strings",
			values: []string{"one", "two", "three"},
			format: "%s",
			want:   []string{"one", "two", "three"},
		},
		{
			name:   "Simple array of integers",
			values: []int{1, 2, 3},
			format: "%d",
			want:   []string{"1", "2", "3"},
		},
		{
			name:   "Simple map of any (int)",
			values: map[string]any{"Tom": 52, "Sue": 48},
			format: "%v",
			want:   []string{"Sue: 48", "Tom: 52"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeArray(tt.values, tt.format)

			if reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("makeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
