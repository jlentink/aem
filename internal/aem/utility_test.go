package aem

import (
	"os/user"
	"reflect"
	"testing"
)

func Test_getCurrentUser(t *testing.T) {
	usr, _ := user.Current()
	mock, _ := user.Current()
	mock.Uid = "0"
	mock.Gid = "0"

	tests := []struct {
		name     string
		want     *user.User
		wantErr  bool
		wantMock bool
	}{
		{
			name:     "get User",
			want:     usr,
			wantErr:  false,
			wantMock: false,
		},
		{
			name:     "get Mock User",
			want:     mock,
			wantErr:  false,
			wantMock: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantMock {
				u = mock
			}
			got, err := getCurrentUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCurrentUser() = %v, want %v", got, tt.want)
			}

			if tt.wantMock && got.Uid != "0" {
				t.Errorf("getCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test https url (low case)",
			args: args{s: "https://www.google.com"},
			want: true,
		},
		{
			name: "Test http url (low case)",
			args: args{s: "http://www.google.com"},
			want: true,
		},
		{
			name: "Test https url (upper case)",
			args: args{s: "HTTPS://WWW.GOOGLE.COM"},
			want: true,
		},
		{
			name: "Test http url (upper case)",
			args: args{s: "HTTP://WWW.GOOGLE.COM"},
			want: true,
		},
		{
			name: "Short",
			args: args{s: "HTTP"},
			want: false,
		},
		{
			name: "Check file path",
			args: args{s: "/etc/hosts"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isURL(tt.args.s); got != tt.want {
				t.Errorf("isURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
