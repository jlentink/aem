package aem

import (
	"os/user"
	"testing"
)

func TestGetJar(t *testing.T) {
	type args struct {
		forceDownload bool
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got, err := GetJar(tt.args.forceDownload, )
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("GetJar() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if got != tt.want {
			//	t.Errorf("GetJar() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestAllowUserStart(t *testing.T) {
	type args struct {
		allow bool
	}
	tests := []struct {
		name string
		args args
		want bool
		uid  string
	}{
		{
			name: "Always allow",
			args: args{allow: true},
			want: true,
			uid:  "501",
		},
		{
			name: "Block root",
			args: args{allow: false},
			want: false,
			uid:  "0",
		},
		{
			name: "Allow normal user",
			args: args{allow: false},
			want: true,
			uid:  "501",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ = user.Current()
			u.Uid = tt.uid
			if got := AllowUserStart(tt.args.allow); got != tt.want {
				t.Errorf("AllowUserStart() = %v, want %v", got, tt.want)
			}
			u = nil
		})
	}
}
