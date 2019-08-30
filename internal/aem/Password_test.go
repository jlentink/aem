package aem

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"runtime"
	"testing"
)

func Test_serviceName(t *testing.T) {
	type args struct {
		i objects.Instance
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Construct string",
			args: args{i: objects.Instance{Name: "Name", Hostname: "Hostname"}},
			want: serviceNameID + "-Name-Hostname",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serviceName(tt.args.i); got != tt.want {
				t.Errorf("serviceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyRingSetPassword(t *testing.T) {
	type args struct {
		i        objects.Instance
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Get password",
			args:    args{i: objects.Instance{Name: "Name", Hostname: "Hostname"}, password: "password"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if runtime.GOOS == "darwin" {
				if err := KeyRingSetPassword(tt.args.i, tt.args.password); (err != nil) != tt.wantErr {
					t.Errorf("KeyRingSetPassword() error = %v, wantErr %v", err, tt.wantErr)
				}

				password, _ := KeyRingGetPassword(tt.args.i)
				if tt.args.password != password {
					t.Errorf("KeyRingSetPassword() error = %s, want %s", password, tt.args.password)
				}

			} else {
				t.Logf("Wrong os skipping. %s", runtime.GOOS)
			}
		})
	}
}

func TestKeyRingGetPassword(t *testing.T) {
	type args struct {
		i objects.Instance
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeyRingGetPassword(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyRingGetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("KeyRingGetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPasswordForInstance(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPasswordForInstance(tt.args.i, tt.args.useKeyring)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPasswordForInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPasswordForInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
