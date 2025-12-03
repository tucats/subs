package subs

import "testing"

func Test_handleSubMap(t *testing.T) {
	tests := []struct {
		name string
		text string
		subs map[string]any
		want string
	}{
		{
			name: "floating format",
			text: "Garbage collection pct of cpu:     {{cpu|%8.7f}}",
			subs: map[string]any{
				"cpu": 1.2,
			},
			want: "Garbage collection pct of cpu:     1.2000000",
		},
		{
			name: "using pad",
			text: `{{addr|%4d}}: {{depth|pad "| "}}{{op}} {{operand}}`,
			subs: map[string]any{
				"addr":    1234,
				"depth":   3,
				"op":      "LOAD_FAST",
				"operand": "42",
			},
			want: `1234: | | | LOAD_FAST 42`,
		},
		{
			name: "complex case 1",
			text: `{{method}} {{endpoint}} {{file}} {{admin|empty|nonempty admin}} {{auth|empty|nonempty auth}}{{perms|label permissions=}}`,
			subs: map[string]any{
				"method":   "GET",
				"endpoint": "/service/proc/Accounts",
				"file":     "accounts.go",
				"admin":    true,
				"auth":     false,
				"perms":    []string{"read", "write"},
			},
			want: `GET /service/proc/Accounts accounts.go admin permissions=[read write]`,
		},
		{
			name: "complex case 2",
			text: `{{method}} {{endpoint}} {{file}} {{admin|empty|nonempty admin}}{{auth|empty|nonempty auth}}{{perms|label permissions=}}`,
			subs: map[string]any{
				"method":   "GET",
				"endpoint": "/service/proc/Accounts",
				"file":     "accounts.go",
				"admin":    false,
				"auth":     true,
				"perms":    []string{},
			},
			want: `GET /service/proc/Accounts accounts.go auth`,
		},
		{
			name: "no subs",
			text: "this is a test",
			want: "this is a test",
		},
		{
			name: "one sub",
			text: "this is a {{kind}} string",
			subs: map[string]any{"kind": "test"},
			want: "this is a test string",
		},
		{
			name: "multiple subs",
			text: "this is a {{kind}} {{item}}",
			subs: map[string]any{
				"kind": "test",
				"item": "value"},
			want: "this is a test value",
		},
		{
			name: "missing sub",
			text: "this is a {{kind}} {{item}}",
			subs: map[string]any{
				"item": "value"},
			want: "this is a !kind! value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleSubstitutionMap(tt.text, tt.subs); got != tt.want {
				t.Errorf("handleSubMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
