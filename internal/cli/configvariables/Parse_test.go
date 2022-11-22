package configvariables

import (
	"reflect"
	"testing"
)

func Test_getVariables(t *testing.T) {
	type args struct {
		content string
	}
	type wants struct {
		variables []string
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name:  "single variable",
			args:  args{content: "Some variable ${variable}"},
			wants: wants{variables: []string{"${variable}"}},
		},
		{
			name:  "single variable",
			args:  args{content: "Some variable ${var1} ${var2~value}"},
			wants: wants{variables: []string{"${var1}", "${var2~value}"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getVariables(tt.args.content)
			if !reflect.DeepEqual(result, tt.wants.variables) {
				t.Errorf("getVariables() = %v, want %v", result, tt.wants.variables)
			}
		})
	}
}

func TestReplaceVariables(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test variables",
			args: args{content: "test ${env:SOMENAME123~123}"},
			want: "test 123",
		},
		{
			name: "Test variables double",
			args: args{content: "test ${env:SOMENAME123~123} ${env:SOMENAME123~123}"},
			want: "test 123 123",
		},
		{
			name: "Test variables",
			args: args{content: "test ${env:SOMENAME123~123} ${env:SOMENAME456~456}"},
			want: "test 123 456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceVariables(tt.args.content); got != tt.want {
				t.Errorf("ReplaceVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
