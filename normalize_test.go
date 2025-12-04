package subs

import (
	"math"
	"reflect"
	"testing"
)

func Test_normalizeNumericValues(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want any
	}{
		{
			name: "float32",
			arg:  float32(123.456),
			want: float64(float32(123.456)),
		},
		{
			name: "float64",
			arg:  1.1,
			want: 1.1,
		},
		{
			name: "float64 to int",
			arg:  float64(-5.0),
			want: int(-5),
		},
		{
			name: "float64 to int overflow",
			arg:  float64(math.MaxInt) + 100.0,
			want: int(math.MaxInt),
		},
		{
			name: "float32 to int",
			arg:  float32(123.000),
			want: int(123),
		},
		{
			name: "float32 to int overflow",
			arg:  float32(math.MaxInt) + 100.0,
			want: int(math.MaxInt),
		},
		{
			name: "int",
			arg:  123,
			want: 123,
		},
		{
			name: "int32",
			arg:  int32(123),
			want: int(123),
		},
		{
			name: "string",
			arg:  "123",
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeNumericValues(tt.arg, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeNumericValues() = %v (%T), want %v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
