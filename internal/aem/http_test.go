package aem

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"runtime"
	"testing"
)

func TestURL(t *testing.T) {
	type args struct {
		i objects.Instance
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test http",
			args: args{i: objects.Instance{
				Protocol: "http",
				Port:     4502,
				Hostname: "some-domain",
				Username: "username",
				Password: "password",
			}},
			want: "http://some-domain:4502",
		},
		{
			name: "Test https",
			args: args{i: objects.Instance{
				Protocol: "https",
				Port:     4503,
				Hostname: "some-domain",
				Username: "username",
				Password: "password",
			}},
			want: "https://some-domain:4503",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLString(&tt.args.i); got != tt.want {
				t.Errorf("URLString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordURL(t *testing.T) {
	type args struct {
		i          objects.Instance
		useKeyring bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test https",
			args: args{i: objects.Instance{
				Name:     "some-host",
				Protocol: "https",
				Port:     4503,
				Hostname: "some-domain",
				Username: "username",
				Password: "password",
			}, useKeyring: false},
			want: "https://username:password@some-domain:4503",
		},
		{
			name: "Test escape",
			args: args{i: objects.Instance{
				Name:     "some-host",
				Protocol: "https",
				Port:     4503,
				Hostname: "some-domain",
				Username: "user:name",
				Password: "pass@word",
			}, useKeyring: false},
			want: "https://user%3Aname:pass%40word@some-domain:4503",
		},
		{
			name: "Test https",
			args: args{i: objects.Instance{
				Name:     "some-host",
				Protocol: "https",
				Port:     4503,
				Hostname: "some-domain",
				Username: "username",
				Password: "password",
			}, useKeyring: true},
			want: "https://username:password@some-domain:4503",
		},
		{
			name: "Test escape",
			args: args{i: objects.Instance{
				Name:     "some-host",
				Protocol: "https",
				Port:     4503,
				Hostname: "some-domain",
				Username: "user:name",
				Password: "pass@word",
			}, useKeyring: true},
			want: "https://user%3Aname:pass%40word@some-domain:4503",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if runtime.GOOS != "darwin" {
				t.Logf("detected non-darwin (%s) disabling keychain test", runtime.GOOS)
				tt.args.useKeyring = false
			} else {
				if tt.args.useKeyring {
					KeyRingSetPassword(tt.args.i, tt.args.i.Password)
				}
			}
			got, err := PasswordURL(tt.args.i, tt.args.useKeyring)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PasswordURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
