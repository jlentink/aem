package objects

import (
	"github.com/jlentink/aem/internal/cli/project"
	"testing"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		want     bool
		wantFile bool
	}{
		{
			name:     "Test if config exists (not exists)",
			want:     false,
			wantFile: false,
		},
		{
			name:     "Test if config exists (exists)",
			want:     true,
			wantFile: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			project.Mock(false)
			if tt.wantFile {
				cf, _ := project.GetConfigFileLocation()
				project.WriteTextFile(cf, "123")
			}

			if got := Exists(); got != tt.want {
				t.Errorf("ConfigExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinStrings(t *testing.T) {
	type args struct {
		s []string
	}

	tests := []struct {
		name string
		want string
		args args
	}{
		{
			name: "Empty slice",
			want: "",
			args: args{s: []string{}},
		},
		{
			name: "single element slice",
			want: "\t\"1\",\n",
			args: args{s: []string{"1"}},
		},
		{
			name: "two element slice",
			want: "\t\"1\",\n\t\"2\",\n",
			args: args{s: []string{"1", "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := joinStrings(tt.args.s)
			if got != tt.want {
				t.Errorf("joinStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Render template",
			want: 2914,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Render()
			if len(got) != tt.want {
				t.Errorf("Render() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestWriteConfigFile(t *testing.T) {
	tests := []struct {
		name      string
		want      int
		wantErr   bool
		wantRO    bool
		wantFsErr bool
	}{
		{
			name:      "Write template",
			want:      2914,
			wantErr:   false,
			wantRO:    false,
			wantFsErr: false,
		},
		{
			name:      "Write template error",
			want:      0,
			wantErr:   true,
			wantRO:    true,
			wantFsErr: false,
		},
		{
			name:      "Write template error",
			want:      0,
			wantErr:   true,
			wantRO:    true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			project.Mock(tt.wantRO)
			project.FilesystemError(tt.wantFsErr)
			written, err := WriteConfigFile()

			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.writeTextFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if written != tt.want {
				t.Errorf("Render() = %v, want %v", written, tt.want)
			}
		})
	}
}
