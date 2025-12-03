package subs

import "testing"

func Test_barEscape(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "quoted sub",
			arg:  `one "|" two`,
			want: `one "!BAR!" two`,
		},
		{
			name: "multiple quoted sub",
			arg:  `one "|" two "|"`,
			want: `one "!BAR!" two "!BAR!"`,
		},
		{
			name: "mixed quoted sub",
			arg:  `one "|" two | three`,
			want: `one "!BAR!" two | three`,
		},
		{
			name: "no sub",
			arg:  "this is a simple string",
			want: "this is a simple string",
		},
		{
			name: "single sub",
			arg:  "one|two",
			want: "one|two",
		},
		{
			name: "multiple sub",
			arg:  "one|two|three",
			want: "one|two|three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := barEscape(tt.arg); got != tt.want {
				t.Errorf("barEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
