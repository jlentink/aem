package casetypes

import (
	"reflect"
	"testing"
)

func TestCamelCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "fooBar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "fooBar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "fooBar0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "fooBar",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "fooBar",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "fooBar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "fooBar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.input); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKababCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "foo-bar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "foo-bar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "foo-bar-0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "foo-bar",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "foo-bar",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "foo-bar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "foo-bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KababCase(tt.args.input); got != tt.want {
				t.Errorf("KababCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "foo bar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "foo bar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "foo bar 0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "foo bar .",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "foo bar ]",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "foo.bar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "foo]bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCase(tt.args.input); got != tt.want {
				t.Errorf("LowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "Foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "FooBar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "Foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "FooBar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "FooBar0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "FooBar",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "FooBar",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "FooBar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "FooBar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.args.input); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "foo_bar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "foo_bar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "foo_bar_0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "foo_bar",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "foo_bar",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "foo_bar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "foo_bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.input); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "FOO",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "FOO BAR",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "FOO",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "FOO BAR",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "FOO BAR 0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "FOO BAR .",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "FOO BAR ]",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "FOO.BAR",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "FOO]BAR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCase(tt.args.input); got != tt.want {
				t.Errorf("UpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: []string{"foo"},
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: []string{"foo", "bar"},
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: []string{"FOO"},
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: []string{"FOO", "BAR"},
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: []string{"foo", "bar", "0"},
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: []string{"foo", "bar", "."},
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: []string{"foo", "bar", "]"},
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: []string{"foo.bar"},
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: []string{"foo]bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenize(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "foo bar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "FOO",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "FOO BAR",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "foo bar 0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "foo bar",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "foo bar",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "foo bar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "foo bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.input); got != tt.want {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitleCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single word",
			args: args{input: "foo"},
			want: "Foo",
		},
		{
			name: "Two words",
			args: args{input: "foo bar"},
			want: "Foo bar",
		},
		{
			name: "Single word (CAPS)",
			args: args{input: "FOO"},
			want: "Foo",
		},
		{
			name: "Two words (CAPS)",
			args: args{input: "FOO BAR"},
			want: "Foo bar",
		},
		{
			name: "Two words with Numeric",
			args: args{input: "foo bar 0"},
			want: "Foo bar 0",
		},
		{
			name: "Two words with Dot",
			args: args{input: "foo bar ."},
			want: "Foo bar .",
		},
		{
			name: "Two words with ]",
			args: args{input: "foo bar ]"},
			want: "Foo bar ]",
		},
		{
			name: "Two words with . in between",
			args: args{input: "foo.bar"},
			want: "Foo.bar",
		},
		{
			name: "Two words with ] in between",
			args: args{input: "foo]bar"},
			want: "Foo]bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TitleCase(tt.args.input); got != tt.want {
				t.Errorf("TitleCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
