package project

import (
	"os"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestMock(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		readOnly bool
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   afero.Fs
	}{
		{
			name: "Get mock Structure object",
			want: afero.NewMemMapFs(),
			args: args{readOnly: false},
		},
		{
			name: "Get readonly mock Structure object",
			want: afero.NewReadOnlyFs(afero.NewMemMapFs()),
			args: args{readOnly: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			got := Mock(tt.args.readOnly)
			if got.Name() != tt.want.Name() {
				t.Errorf("Mock() = %v, want %v", got.Name(), tt.want.Name())
			}
		})
	}
}

func TestFilesystemError(t *testing.T) {
	type args struct {
		err bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Set filesystem error",
			args: args{err: true},
			want: true,
		}, {
			name: "Set filesystem error",
			args: args{err: false},
			want: false,
		},
	}
	for _, tt := range tests {
		setUpTest()
		t.Run(tt.name, func(t *testing.T) {
			FilesystemError(tt.args.err)
			if tt.want != fsErr {
				t.Errorf("FilesystemError() = %v, want %v", fsErr, tt.want)
			}
		})
	}
}

func Test_filesystem(t *testing.T) {
	tests := []struct {
		name string
		want afero.Fs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filesystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filesystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendSlash(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "With slash",
			args: args{path: "/a/"},
			want: "/a/",
		},
		{
			name: "Without slash",
			args: args{path: "/a"},
			want: "/a/",
		},
		{
			name: "Single slash",
			args: args{path: "/"},
			want: "/",
		},
		{
			name: "Empty",
			args: args{path: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			if got := appendSlash(tt.args.path); got != tt.want {
				t.Errorf("Structure.appendSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSlash(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "with slash",
			args: args{path: "/a/"},
			want: "/a",
		},
		{
			name: "without slash",
			args: args{path: "/a"},
			want: "/a",
		},
		{
			name: "single slash",
			args: args{path: "/"},
			want: "",
		},
		{
			name: "empty",
			args: args{path: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			if got := RemoveSlash(tt.args.path); got != tt.want {
				t.Errorf("Structure.removeSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeString(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Normalize string 1",
			args: args{input: "1 2 3"},
			want: "1-2-3",
		},
		{
			name: "Normalize string 2",
			args: args{input: "1#2$3"},
			want: "1-2-3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			if got := normalizeString(tt.args.input); got != tt.want {
				t.Errorf("Structure.normalizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		path string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           bool
		wantCreateFile bool
	}{
		{
			name:           "Check if file exists",
			args:           args{path: "/sample-file"},
			wantCreateFile: true,
			want:           true,
		},
		{
			name:           "Check if file exists",
			args:           args{path: "/sample-file"},
			wantCreateFile: false,
			want:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			Mock(false)
			if tt.wantCreateFile {
				afero.WriteFile(filesystem(), tt.args.path, []byte{}, 0644)
			}
			if got := Exists(tt.args.path); got != tt.want {
				t.Errorf("Structure.exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRename(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantFile  bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:      "Rename file 1",
			args:      args{source: "/some-file-1", destination: "/some-file-2"},
			wantErr:   false,
			wantFile:  true,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Source does not exists",
			args:      args{source: "/some-file-1", destination: "/some-file-2"},
			wantErr:   true,
			wantFile:  false,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Read-only file system",
			args:      args{source: "/some-file-1", destination: "/some-file-2"},
			wantErr:   true,
			wantFile:  false,
			wantFsErr: false,
			wantRO:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			Mock(tt.wantRO)
			if tt.wantFile {
				afero.WriteFile(filesystem(), tt.args.source, []byte{}, 0644)
			}
			if err := Rename(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("Structure.rename() error = %v, wantErr %v", err, tt.wantErr)
			}
			if Exists(tt.args.source) && !tt.wantErr {
				t.Errorf("Structure.rename() source file still exists")
			}
			if !Exists(tt.args.destination) && !tt.wantErr {
				t.Errorf("Structure.rename() destination file does not exists")
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		src     string
		dst     string
		content string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      int64
		wantErr   bool
		wantFile  bool
		wantDir   bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:     "Copy file",
			args:     args{src: "/some-file-1", dst: "/some-file-2", content: "123"},
			want:     3,
			wantErr:  false,
			wantFile: true,
			wantDir:  false,
			wantRO:   false,
		},
		{
			name:     "Copy file error src",
			args:     args{src: "/some-file-1", dst: "/some-file-2", content: "123"},
			want:     0,
			wantErr:  true,
			wantFile: false,
			wantDir:  false,
			wantRO:   false,
		},
		{
			name:     "Copy file error ro",
			args:     args{src: "/some-file-1", dst: "/some-file-2", content: "123"},
			want:     0,
			wantErr:  true,
			wantFile: false,
			wantDir:  false,
			wantRO:   true,
		},
		{
			name:     "Copy dir error",
			args:     args{src: "/some-file-1", dst: "/some-file-2", content: "123"},
			want:     0,
			wantErr:  true,
			wantFile: false,
			wantDir:  true,
			wantRO:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			Mock(tt.wantRO)
			if tt.wantFile {
				afero.WriteFile(filesystem(), tt.args.src, []byte(tt.args.content), 0644)
			}
			if tt.wantDir {
				filesystem().MkdirAll(tt.args.src, 0755)
			}

			got, err := Copy(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.copy() = %v, want %v", got, tt.want)
			}
			if !Exists(tt.args.dst) && !tt.wantErr {
				t.Errorf("Structure.copy() destination file does not exists")
			}

		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    afero.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			got, err := Create(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		want        []os.FileInfo
		wantErr     bool
		wantMock    bool
		wantMockLen int
	}{
		{
			name:        "Read simple dir",
			args:        args{path: "/test/dir/"},
			wantErr:     false,
			want:        nil,
			wantMock:    true,
			wantMockLen: 3,
		},
		{
			name:        "Dir not exists",
			args:        args{path: "/somerandom-folder/"},
			wantErr:     true,
			want:        nil,
			wantMock:    false,
			wantMockLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			Mock(false)
			if tt.wantMock {
				err := fs.MkdirAll(tt.args.path, 0777)
				if err != nil {
					t.Errorf("ReadDir() Could not create mock dir")
				}
				for i := 0; i < tt.wantMockLen; i++ {
					WriteTextFile(tt.args.path+string(i)+".txt", string(i))
				}
			}
			got, err := ReadDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantMockLen {
				t.Errorf("ReadDir() = %v, want %v", len(got), tt.wantMockLen)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		want       []os.FileInfo
		wantErr    bool
		wantMockRO bool
		wantFile   bool
		wantFsErr  bool
	}{
		{
			name:       "Delete file",
			args:       args{path: "/test.txt"},
			wantErr:    false,
			want:       nil,
			wantMockRO: false,
			wantFile:   true,
		},
		{
			name:       "Delete file RO",
			args:       args{path: "/test.txt"},
			wantErr:    true,
			want:       nil,
			wantMockRO: true,
			wantFile:   true,
		},
		{
			name:       "Delete file that not exists",
			args:       args{path: "/test.txt"},
			wantErr:    true,
			want:       nil,
			wantMockRO: false,
			wantFile:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			Mock(tt.wantMockRO)
			FilesystemError(tt.wantFsErr)
			if tt.wantFile {
				WriteTextFile(tt.args.path, "123")
			}

			got := Remove(tt.args.path)
			if tt.wantErr && got == nil {
				t.Errorf("Remove() = %v, want %v", got, tt.wantErr)
			}

		})
	}

}
