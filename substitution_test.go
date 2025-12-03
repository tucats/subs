package subs

import (
	"testing"
)

func TestSubstitution(t *testing.T) {
	type D1 struct {
		Name string
	}

	type D2 struct {
		User  D1
		Count int
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text   string
		values []any
		want   string
	}{
		{
			name: "struct nested member with formats",
			text: "There {{Count|card is,are}} {{Count}} {{Count|card user,users}}",
			values: []any{
				D2{
					User: D1{
						Name: "Sam",
					},
					Count: 3},
			},
			want: "There are 3 users",
		},
		{
			name: "struct nested member",
			text: "The user is {{User.Name}}",
			values: []any{
				D2{
					User: D1{
						Name: "Sam",
					},
					Count: 55},
			},
			want: "The user is Sam",
		},
		{
			name: "struct member",
			text: "The count is {{Name}}",
			values: []any{
				D1{Name: "42"},
			},
			want: "The count is 42",
		},
		{
			name: "map element",
			text: "The count is {{name}}",
			values: []any{
				map[string]string{"name": "42"},
			},
			want: "The count is 42",
		},
		{
			name:   "single item",
			text:   "The count is {{.}}",
			values: []any{42},
			want:   "The count is 42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Substitution(tt.text, tt.values...)
			if got != tt.want {
				t.Errorf("Substitution() = %v, want %v", got, tt.want)
			}
		})
	}
}
