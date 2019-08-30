package project

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/spf13/afero"
	"os"
	"testing"
)

func Test_createBinDir(t *testing.T) {
	dir, _ := os.Getwd()
	type fields struct {
		fs afero.Fs
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantErr   bool
		wantRO    bool
		wantFSErr bool
	}{
		{
			name:      "Create bin dir",
			want:      dir + `/` + configAemInstancesDir + `/` + configBinDir + `/`,
			wantErr:   false,
			wantRO:    false,
			wantFSErr: false,
		},
		{
			name:      "Create bin dir err 1",
			want:      ``,
			wantErr:   true,
			wantRO:    true,
			wantFSErr: false,
		},
		{
			name:      "Create bin dir err 2",
			want:      ``,
			wantErr:   true,
			wantRO:    false,
			wantFSErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRO)
			FilesystemError(tt.wantFSErr)
			got, err := CreateBinDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.createBinDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.createBinDir() = %v, want %v", got, tt.want)
			}

			if exists, _ := afero.Exists(filesystem(), tt.want); !exists && !tt.wantErr {
				t.Errorf("Structure.createBinDir() no directory rreation, %s", tt.want)
			}
		})
	}
}

func TestCreateInstanceDir(t *testing.T) {
	dir, _ := os.Getwd()
	i := objects.Instance{Name: "1 2 3"}

	type fields struct {
		fs afero.Fs
	}
	type args struct {
		instance objects.Instance
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantErr   bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:      "Create aem instance dir",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3/",
			wantErr:   false,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Create aem instance dir fs error 1",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
			wantRO:    false,
		},
		{
			name:      "Create aem instance dir fs error 2",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: false,
			wantRO:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRO)
			FilesystemError(tt.wantFsErr)
			got, err := CreateInstanceDir(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.CreateInstanceDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.CreateInstanceDir() = %v, want %v", got, tt.want)
				return
			}
			if exists, _ := afero.Exists(filesystem(), tt.want); !exists && !tt.wantErr {
				t.Errorf("Structure.CreateInstanceDir() no directory creation, %s", tt.want)
			}
		})
	}
}

func Test_createAemInstallDir(t *testing.T) {
	dir, _ := os.Getwd()
	i := objects.Instance{Name: "1 2 3"}

	type fields struct {
		fs afero.Fs
	}
	type args struct {
		instance objects.Instance
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantErr   bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:      "Create aem install dir",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3/" + configAemRunDir + "/" + configAemInstallDir,
			wantErr:   false,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Create aem install dir fs error 1",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
			wantRO:    false,
		},
		{
			name:      "Create aem install dir fs error 2",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: false,
			wantRO:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRO)
			FilesystemError(tt.wantFsErr)
			got, err := createAemInstallDir(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.createAemInstallDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.createAemInstallDir() = %v, want %v", got, tt.want)
				return
			}
			if exists, _ := afero.Exists(filesystem(), tt.want); !exists && !tt.wantErr {
				t.Errorf("Structure.createAemInstallDir() no directory creation, %s", tt.want)
			}
		})
	}
}

func TestCreateInstancesDir(t *testing.T) {
	dir, _ := os.Getwd()
	type fields struct {
		fs afero.Fs
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantErr   bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:      "Create instances dir",
			want:      dir + "/" + configAemInstancesDir + "/",
			wantErr:   false,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Create instance dir fs error 1",
			want:      ``,
			wantErr:   true,
			wantFsErr: false,
			wantRO:    true,
		},
		{
			name:      "Create instance dir fs error 2",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
			wantRO:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRO)
			FilesystemError(tt.wantFsErr)
			got, err := CreateInstancesDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.createInstancesDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.createInstancesDir() = %v, want %v", got, tt.want)
				return
			}
			if exists, _ := afero.Exists(filesystem(), tt.want); !exists && !tt.wantErr {
				t.Errorf("Structure.createInstancesDir() no directory creation, %s", tt.want)
			}
		})
	}
}

func Test_createDirForPackage(t *testing.T) {
	dir, _ := os.Getwd()
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		aemPackage objects.Package
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantErr   bool
		wantFsErr bool
		wantRO    bool
	}{
		{
			name:      "Create pkg dir",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0"}},
			want:      dir + "/" + configAemInstancesDir + "/" + configPackageDir + "/some-name/1.0.0",
			wantErr:   false,
			wantFsErr: false,
			wantRO:    false,
		},
		{
			name:      "Create pkg dir error 1",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0"}},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
			wantRO:    false,
		},
		{
			name:      "Create pkg dir error 2",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0"}},
			want:      ``,
			wantErr:   true,
			wantFsErr: false,
			wantRO:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock(tt.wantRO)
			FilesystemError(tt.wantFsErr)
			got, err := CreateDirForPackage(&tt.args.aemPackage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.createDirForPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.createDirForPackage() = %v, want %v", got, tt.want)
			}
			if exists, _ := afero.Exists(filesystem(), tt.want); !exists && !tt.wantErr {
				t.Errorf("Structure.createDirForPackage() no directory creation, %s", tt.want)
			}
		})
	}
}
