package version

import "testing"

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test version return",
			want: "1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Main = tt.want
			if got := GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBuild(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test build return",
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Build = tt.want
			if got := GetBuild(); got != tt.want {
				t.Errorf("GetBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisplayVersion(t *testing.T) {
	type args struct {
		v bool
		m bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		build   string
		version string
	}{
		{
			name:    "Get normal version",
			args:    args{v: false, m: false},
			version: "1.2.3",
			build:   "1234567890",
			want:    "AEMcli\nVersion: 1.2.3\n",
		},
		{
			name:    "Get Minimal version",
			args:    args{v: false, m: true},
			version: "1.2.3",
			build:   "1234567890",
			want:    "1.2.3\n",
		},
		{
			name:    "Get Verbose version",
			args:    args{v: true, m: false},
			version: "1.2.3",
			build:   "1234567890",
			want:    "AEMcli (https://github.com/jlentink/aem)\nVersion: 1.2.3\nBuilt: 1234567890\n",
		},
		{
			name:    "Get Verbose & Minal version",
			args:    args{v: true, m: true},
			version: "1.2.3",
			build:   "1234567890",
			want:    "1.2.3\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Build = tt.build
			Main = tt.version
			if got := DisplayVersion(tt.args.v, tt.args.m); got != tt.want {
				t.Errorf("DisplayVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
