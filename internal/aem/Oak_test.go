package aem

import "testing"

func TestOakVersionForAEM(t *testing.T) {
	type args struct {
		v string
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Default",
			args: args{v: "999.9", d: "1.0"},
			want: "1.0",
		},
		{
			name: "Get for AEM 6.0",
			args: args{v: "6.0", d: "1.0"},
			want: "1.0.42",
		},
		{
			name: "Get for AEM 6.1",
			args: args{v: "6.1", d: "1.0"},
			want: "1.2.31",
		},
		{
			name: "Get for AEM 6.2",
			args: args{v: "6.2", d: "1.0"},
			want: "1.4.24",
		},
		{
			name: "Get for AEM 6.3",
			args: args{v: "6.3", d: "1.0"},
			want: "1.6.16",
		},
		{
			name: "Get for AEM 6.4",
			args: args{v: "6.4", d: "1.0"},
			want: "1.8.12",
		},
		{
			name: "Get for AEM 6.5",
			args: args{v: "6.5", d: "1.0"},
			want: "1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oakVersionForAEM(tt.args.v, tt.args.d); got != tt.want {
				t.Errorf("OakVersionForAEM() = %v, want %v", got, tt.want)
			}
		})
	}
}
