package http

import (
	"net/url"
	"testing"
)

func Test_URLToUrlString(t *testing.T) {
	type args struct {
		u *url.URL
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "plain",
			args: args{s: "http://www.google.com"},
			want: "http://www.google.com/",
		},
		{
			name: "auth",
			args: args{s: "http://1:2@www.google.com"},
			want: "http://www.google.com/",
		},
		{
			name: "query",
			args: args{s: "http://www.google.com?a=b"},
			want: "http://www.google.com/?a=b",
		},
		{
			name: "path",
			args: args{s: "http://www.google.com/a/b.html"},
			want: "http://www.google.com/a/b.html",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.u, _ = url.Parse(tt.args.s)
			if got := URLToURLString(tt.args.u); got != tt.want {
				t.Errorf("urlToUrlString() = %v, want %v", got, tt.want)
			}
		})
	}
}
