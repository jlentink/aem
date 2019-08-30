package project

import "testing"

func TestWriteTextFile(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		wantFile bool
		wantRo   bool
		wantErr  bool
	}{
		{
			name:     "Write file",
			wantFile: true,
			wantRo:   false,
			wantErr:  false,
		},
		{
			name:     "Get RO error",
			want:     true,
			wantFile: true,
			wantRo:   true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRo)
			_, err := WriteTextFile(tt.args.path, tt.args.content)

			if tt.wantErr && err == nil {
				t.Errorf("WriteTextFile() = %v, want err %v", err, tt.want)
			}

			if got := Exists(tt.args.path); got != tt.wantFile {
				t.Errorf("WriteTextFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteGitIgnoreFile(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		wantFile  bool
		wantFSErr bool
		wantFSRO  bool
	}{
		{
			name:      "Test if ignore exists",
			want:      true,
			wantFile:  true,
			wantFSErr: false,
			wantFSRO:  false,
		},
		{
			name:      "FS err",
			want:      true,
			wantFile:  true,
			wantFSErr: true,
			wantFSRO:  false,
		},
		{
			name:      "RO err",
			want:      false,
			wantFile:  false,
			wantFSErr: false,
			wantFSRO:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantFSRO)
			FilesystemError(tt.wantFSErr)
			path, _ := getIgnoreFileLocation()
			WriteGitIgnoreFile()

			if got := Exists(path); got != tt.want {
				t.Errorf("WriteGitIgnoreFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
