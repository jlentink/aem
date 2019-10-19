package generate

import (
	"testing"
)

func Test_ParseTemplate(t *testing.T) {
	type args struct {
		tpl  string
		vars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test simple template",
			args:    args{tpl: "{{.Hello}} world.", vars: map[string]string{"Hello": "Hello"}},
			want:    "Hello world.",
			wantErr: false,
		},
		{
			name:    "Test uppercase template",
			args:    args{tpl: "{{UC .Hello}} world.", vars: map[string]string{"Hello": "Hello"}},
			want:    "HELLO world.",
			wantErr: false,
		},
		{
			name:    "Test uppercase template",
			args:    args{tpl: "{{LC .Hello}} world.", vars: map[string]string{"Hello": "Hello"}},
			want:    "hello world.",
			wantErr: false,
		},
		{
			name:    "Test uppercase template",
			args:    args{tpl: "{{BLA .Hello}} world.", vars: map[string]string{"Hello": "Hello"}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplate(tt.args.tpl, tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
