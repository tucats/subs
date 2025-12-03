package subs

import "testing"

func Test_handleSub(t *testing.T) {
	tests := []struct {
		name string
		text string
		subs map[string]any
		want string
	}{
		{
			name: "short cardinal with value 0",
			text: `{{value|cardinal frog, frogs}}`,
			subs: map[string]any{"value": 0},
			want: "frogs",
		},
		{
			name: "short cardinal with value 1",
			text: `{{value|cardinal frog, frogs}}`,
			subs: map[string]any{"value": 1},
			want: "frog",
		},
		{
			name: "short cardinal with value 7",
			text: `{{value|cardinal frog, frogs}}`,
			subs: map[string]any{"value": 7},
			want: "frogs",
		},
		{
			name: "cardinal with value 0",
			text: `{{value|cardinal "frog free", frog, frogs}}`,
			subs: map[string]any{"value": 0},
			want: "frog free",
		},
		{
			name: "cardinal with value 1",
			text: `{{value|card "frog free", frog, frogs}}`,
			subs: map[string]any{"value": 1},
			want: "frog",
		},
		{
			name: "cardinal with value 33",
			text: `{{value|card "frog free", frog, frogs}}`,
			subs: map[string]any{"value": 33},
			want: "frogs",
		},
		{
			name: "magnitude with value 0",
			text: `{{value|zero "frog free"|one frog|many frogs}}`,
			subs: map[string]any{"value": 0},
			want: "frog free",
		},
		{
			name: "magnitude with value 1",
			text: `{{value|zero "frog free"|one frog|many frogs}}`,
			subs: map[string]any{"value": 1},
			want: "frog",
		},
		{
			name: "magnitude with value 12",
			text: `{{value|zero "frog free"|one frog|many frogs}}`,
			subs: map[string]any{"value": 12},
			want: "frogs",
		},
		{
			name: "left",
			text: "{{value|left 8}}",
			subs: map[string]any{"value": "abc"},
			want: "abc     ",
		},
		{
			name: "center",
			text: "{{value|center 8}}",
			subs: map[string]any{"value": "abc"},
			want: "  abc   ",
		},
		{
			name: "right",
			text: "{{value|right 8}}",
			subs: map[string]any{"value": "abc"},
			want: "     abc",
		},
		{
			name: "format with center",
			text: "{{value|%3.1f|center 8}}",
			subs: map[string]any{"value": 5.6},
			want: "  5.6   ",
		},
		{
			name: "combo format and list",
			text: "{{value|%02d|list}}",
			subs: map[string]any{"value": []any{1, 2, 3}},
			want: "01, 02, 03",
		},
		{
			name: "combo format, list, and size",
			text: "{{value|%02d|list|size 8}}",
			subs: map[string]any{"value": []any{1, 2, 3}},
			want: "01, 0...",
		},
		{
			name: "combo list and size truncated",
			text: "{{value|list|size 10}}",
			subs: map[string]any{"value": []string{"one", "two", "three"}},
			want: "one, tw...",
		},
		{
			name: "combo list and size not truncated",
			text: "{{value|list|size 20}}",
			subs: map[string]any{"value": []string{"one", "two", "three"}},
			want: "one, two, three",
		},
		{
			name: "size not needed",
			text: `{{value|size 10}}`,
			subs: map[string]any{"value": "test"},
			want: "test",
		},
		{
			name: "size needed",
			text: `{{value|size 10}}`,
			subs: map[string]any{"value": "test string of text"},
			want: "test st...",
		},
		{
			name: "size invalid",
			text: `{{value|size 2}}`,
			subs: map[string]any{"value": "test string of text"},
			want: "!Invalid size: 2!",
		},
		{
			name: "simple pad",
			text: `{{size|pad "*"}}`,
			subs: map[string]any{"size": 3},
			want: "***",
		},
		{
			name: "complex pad",
			text: `{{size|pad "XO"}}`,
			subs: map[string]any{"size": 2},
			want: "XOXO",
		},
		{
			name: "zero pad",
			text: `{{size|pad "*"}}`,
			subs: map[string]any{"size": 0},
			want: "",
		},
		{
			name: "label zero value",
			text: `{{item|label "flag="}}`,
			subs: map[string]any{"item": 0},
			want: "",
		},
		{
			name: "explicit format string sub",
			text: "{{item|format %10s}}",
			subs: map[string]any{"item": "test"},
			want: "      test",
		},
		{
			name: "label zero value with empty value",
			text: `{{item|label "flag="|%05x | empty none}}`,
			subs: map[string]any{"item": 0},
			want: "none",
		},
		{
			name: "label non-zero value with format",
			text: `{{item|label "flag="|%05x}}`,
			subs: map[string]any{"item": 5},
			want: "flag=00005",
		},
		{
			name: "label zero value with format",
			text: `{{item|label "flag="|%05x}}`,
			subs: map[string]any{"item": 0},
			want: "00000",
		},
		{
			name: "label non-zero value",
			text: `{{item|label "flag="}}`,
			subs: map[string]any{"item": 33},
			want: "flag=33",
		},
		{
			name: "left-justify string sub",
			text: "{{item|%-10s}}",
			subs: map[string]any{"item": "test"},
			want: "test      ",
		},
		{
			name: "right-justify string sub",
			text: "{{item|%10s}}",
			subs: map[string]any{"item": "test"},
			want: "      test",
		},
		{
			name: "simple string sub",
			text: "{{item}}",
			subs: map[string]any{"item": "test"},
			want: "test",
		},
		{
			name: "simple boolean sub",
			text: "{{item}}",
			subs: map[string]any{"item": false},
			want: "false",
		},
		{
			name: "simple array sub",
			text: "{{item}}",
			subs: map[string]any{"item": []string{"one", "two"}},
			want: "[one two]",
		},
		{
			name: "no sub",
			text: "This is a simple string",
			want: "This is a simple string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleFormat(tt.text, tt.subs); got != tt.want {
				t.Errorf("handleSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
