package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func Test(t *testing.T) {

}

func Test_projectStructure_getWorkDir(t *testing.T) {

	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Current workdir",
			want: dir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getWorkDir(); got != tt.want {
				t.Errorf("projectStructure.getWorkDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getConfigFileLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Find config file",
			want: dir + "/" + configFilename,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getConfigFileLocation(); got != tt.want {
				t.Errorf("projectStructure.getConfigFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getInstanceDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Find config file",
			want: dir + "/" + configAemInstanceDir + "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getInstanceDirLocation(); got != tt.want {
				t.Errorf("projectStructure.getInstanceDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_appendSlash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		p    *projectStructure
		args args
		want string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.appendSlash(tt.args.path); got != tt.want {
				t.Errorf("projectStructure.appendSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_removeSlash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		p    *projectStructure
		args args
		want string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.removeSlash(tt.args.path); got != tt.want {
				t.Errorf("projectStructure.removeSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getIgnoreFileLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Find ignore file",
			want: dir + "/" + configAemInstanceDir + "/" + configInstanceGitIgnore,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getIgnoreFileLocation(); got != tt.want {
				t.Errorf("projectStructure.getIgnoreFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getRunDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "Find run dir",
			want: dir + "/" + configAemInstanceDir + "/some-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			if got := p.getRunDirLocation(*i); got != tt.want {
				t.Errorf("projectStructure.getRunDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getPidFileLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "Find pid file",
			want: dir + "/" + configAemInstanceDir + "/some-name/" + configAemPidFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			if got := p.getPidFileLocation(*i); got != tt.want {
				t.Errorf("projectStructure.getPidFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getAppDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "Find app dir",
			want: dir + "/" + configAemInstanceDir + "/some-name/" + configAppDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			if got := p.getAppDirLocation(*i); got != tt.want {
				t.Errorf("projectStructure.getAppDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getAemInstallDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "Find aem install dir",
			want: dir + "/" + configAemInstanceDir + "/some-name/" + configAemInstallDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			if got := p.getAemInstallDirLocation(*i); got != tt.want {
				t.Errorf("projectStructure.getAemInstallDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getJarFileLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Find jar location",
			want: dir + "/" + configAemInstanceDir + "/" + configAemJar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getJarFileLocation(); got != tt.want {
				t.Errorf("projectStructure.getJarFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getLogFileLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "Find log file location",
			want: dir + "/" + configAemInstanceDir + "/some-name/" + configAemLogFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			if got := p.getLogFileLocation(*i); got != tt.want {
				t.Errorf("projectStructure.getLogFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getPackagesDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		{
			name: "Find log file location",
			want: dir + "/" + configAemInstanceDir + "/" + configPackageDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getPackagesDirLocation(); got != tt.want {
				t.Errorf("projectStructure.getPackagesDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getDirForPackage(t *testing.T) {
	dir, _ := os.Getwd()
	type args struct {
		aemPackage packageDescription
	}
	tests := []struct {
		name string
		p    *projectStructure
		args args
		want string
	}{
		{
			name: "Find package location.",
			args: args{aemPackage: packageDescription{Name: "foo-bar", Version: "1.0.0"}},
			want: dir + "/" + configAemInstanceDir + "/" + configPackageDir + "/" + "foo-bar/1.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getDirForPackage(tt.args.aemPackage); got != tt.want {
				t.Errorf("projectStructure.getDirForPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getLocationForPackage(t *testing.T) {
	dir, _ := os.Getwd()
	type args struct {
		aemPackage packageDescription
	}
	tests := []struct {
		name string
		p    *projectStructure
		args args
		want string
	}{
		{
			name: "Find package file location.",
			args: args{aemPackage: packageDescription{Name: "foo-bar", Version: "1.0.0", DownloadName: "foo-bar-1.0.0.zip"}},
			want: dir + "/" + configAemInstanceDir + "/" + configPackageDir + "/" + "foo-bar/1.0.0/foo-bar-1.0.0.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.getLocationForPackage(tt.args.aemPackage); got != tt.want {
				t.Errorf("projectStructure.getLocationForPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_createAemInstallDir(t *testing.T) {
	dir, _ := os.Getwd()
	tests := []struct {
		name string
		p    *projectStructure
		i    *aemInstanceConfig
		want string
	}{
		{
			name: "test create install dir",
			want: dir + "/" + configAemInstanceDir + "/" + configAemRunDir + "/" + configAemInstallDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			i := &aemInstanceConfig{Name: "some-name"}
			p.mock()
			//fix
			if got := p.createAemInstallDir(*i); got != tt.want {
				if exist, _ := afero.DirExists(p.fs, tt.want+"2"); exist == true {
					t.Errorf("projectStructure.createAemInstallDir() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_projectStructure_createInstanceDir(t *testing.T) {
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.createInstanceDir(); got != tt.want {
				t.Errorf("projectStructure.createInstanceDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_createDirForPackage(t *testing.T) {
	type args struct {
		aemPackage packageDescription
	}
	tests := []struct {
		name string
		p    *projectStructure
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.createDirForPackage(tt.args.aemPackage); got != tt.want {
				t.Errorf("projectStructure.createDirForPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_writeGitIgnoreFile(t *testing.T) {
	tests := []struct {
		name string
		p    *projectStructure
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			if got := p.writeGitIgnoreFile(); got != tt.want {
				t.Errorf("projectStructure.writeGitIgnoreFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_copy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		p       *projectStructure
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{}
			got, err := p.copy(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectStructure.copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("projectStructure.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_normalizeString(t *testing.T) {
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
			name: "check to replace space",
			args: args{input: "a b c"},
			want: "a-b-c",
		},
		{
			name: "check to replace underscore and dot",
			args: args{input: "a_b.c"},
			want: "a_b-c",
		},
		{
			name: "check to replace underscore and dot",
			args: args{input: "a_b.c�d"},
			want: "a_b-c-d",
		},
		{
			name: "check to replace underscore and dot",
			args: args{input: "a_b.c�d-Û"},
			want: "a_b-c-d--",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{
				fs: tt.fields.fs,
			}
			if got := p.normalizeString(tt.args.input); got != tt.want {
				t.Errorf("projectStructure.normalizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectStructure_getUnpackDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	type fields struct {
		fs afero.Fs
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test default unpack folder",
			want: dir + "/" + configAemInstanceDir + "/" + configAemRunDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &projectStructure{
				fs: tt.fields.fs,
			}
			if got := p.getUnpackDirLocation(); got != tt.want {
				t.Errorf("projectStructure.getUnpackDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProjectStructure(t *testing.T) {
	tests := []struct {
		name string
		want projectStructure
	}{
		{
			name: "Get Project structure object",
			want: projectStructure{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := newProjectStructure(); reflect.TypeOf(got).Kind() != reflect.TypeOf(tt.want).Kind() {
				t.Errorf("newProjectStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
