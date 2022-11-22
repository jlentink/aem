package configvariables

import (
	"reflect"
	"testing"
)

func Test_envVariable_Variable(t *testing.T) {
	type fields struct {
		raw          string
		name         string
		defaultValue string
	}
	type args struct {
		content string
	}

	type wants struct {
		name  string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   wants
	}{
		{
			name: "Simple variable",
			args: args{content: "${env:simple}"},
			want: wants{name: "simple"},
		},
		{
			name: "default variable",
			args: args{content: "${env:simple~default}"},
			want: wants{
				name:  "simple",
				value: "default",
			},
		},
		{
			name: "default variable",
			args: args{content: "${env:simple~}"},
			want: wants{
				name:  "simple",
				value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &envVariable{
				raw:          tt.fields.raw,
				name:         tt.fields.name,
				defaultValue: tt.fields.defaultValue,
			}
			v.Variable(tt.args.content)
			if got := v.Name(); got != tt.want.name {
				t.Errorf("Name() = %v, want %v", got, tt.want.name)
			}
			if tt.want.value != "" {
				if got := v.DefaultValue(); got != tt.want.value {
					t.Errorf("Value() = %v, want %v", got, tt.want.value)
				}

			}
		})
	}
}

func Test_newEnvVariable(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want envVariable
	}{
		{
			name: "Test new env variable",
			args: args{content: "${env:test}"},
			want: envVariable{
				raw:          "env:test",
				name:         "test",
				defaultValue: "",
				varType:      "ENV",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEnvVariable(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEnvVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_envVariable_Value(t *testing.T) {
	type fields struct {
		raw          string
		name         string
		defaultValue string
	}
	type wants struct {
		value string
		exact bool
	}
	tests := []struct {
		name   string
		fields fields
		want   wants
	}{
		{
			name: "Get user",
			fields: fields{
				raw:          "env:USER",
				name:         "USER",
				defaultValue: "",
			},
			want: wants{
				value: "",
				exact: false,
			},
		},
		{
			name: "Get default value",
			fields: fields{
				raw:          "USER-23ii3u23u3io2uoi3uio",
				name:         "USER-23ii3u23u3io2uoi3uio",
				defaultValue: "TEST",
			},
			want: wants{
				value: "TEST",
				exact: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &envVariable{
				raw:          tt.fields.raw,
				name:         tt.fields.name,
				defaultValue: tt.fields.defaultValue,
			}
			if tt.want.exact {
				if got := v.Value(); got != tt.want.value {
					t.Errorf("Value() = %v, want %v", got, tt.want)
				}
			} else {
				if got := v.Value(); len(got) == 0 {
					t.Errorf("Value() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
