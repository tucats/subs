package subs

import (
	"reflect"
	"testing"
)

func Test_splitOutFormats(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{
			name: "sub at end",
			arg:  "string {{item}}",
			want: []string{"string ", "{{item}}"},
		},
		{
			name: "sub at start",
			arg:  "{{item}} string",
			want: []string{"{{item}}", " string"},
		},
		{
			name: "single part",
			arg:  "simple string",
			want: []string{"simple string"},
		},
		{
			name: "one sub",
			arg:  "test {{item}} string",
			want: []string{"test ", "{{item}}", " string"},
		},
		{
			name: "multiple subs",
			arg:  "test {{item1}} and {{item2}} string",
			want: []string{"test ", "{{item1}}", " and ", "{{item2}}", " string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitOutFormats(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
