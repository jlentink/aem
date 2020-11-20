package project

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/spf13/afero"
	"os"
	"testing"
)

func setUpTest() {
	fs = afero.NewOsFs()
	fsErr = false
}

func Test_getWorkDir(t *testing.T) {
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
	}{
		{
			name:      "Get working dir",
			want:      dir,
			wantFsErr: false,
		},
		{
			name:      "Get working dir error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetWorkDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetWorkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetWorkDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBinDir(t *testing.T) {
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
	}{
		{
			name:      "Get bin dir location",
			want:      dir + `/` + configAemInstancesDir + `/` + configBinDir + `/`,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get bin dir location",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetBinDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getBinDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getBinDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfigFileLocation(t *testing.T) {
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
	}{
		{
			name:      "Get Instance file location",
			want:      dir + "/" + configFilename,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get Instance file location error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetConfigFileLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getConfigFileLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getConfigFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInstancesDirLocation(t *testing.T) {
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
	}{
		{
			name:      "Get instance dir location",
			want:      dir + "/" + configAemInstancesDir + "/",
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get instance dir location error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetInstancesDirLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetInstanceDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetInstanceDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLicenseLocation(t *testing.T) {
	dir, _ := os.Getwd()
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
	}{
		{
			name:      "Get license location",
			want:      dir + "/" + configAemInstancesDir + "/123/" + configLicenseFile,
			args:      args{instance: objects.Instance{Name: "123"}},
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get license location error",
			want:      ``,
			args:      args{instance: objects.Instance{Name: "123"}},
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetLicenseLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetLicenseLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetLicenseLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIgnoreFileLocation(t *testing.T) {
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
	}{
		{
			name:      "Find ignore file",
			want:      dir + "/" + configAemInstancesDir + "/" + configInstanceGitIgnore,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Find ignore file",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := getIgnoreFileLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getIgnoreFileLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getIgnoreFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUnpackDirLocation(t *testing.T) {
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
	}{
		{
			name:      "Find unpack dir location",
			want:      dir + `/` + configAemInstancesDir + "/" + configAemRunDir,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Find unpack dir location error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetUnpackDirLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getUnpackDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getUnpackDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInstanceDirLocation(t *testing.T) {
	dir, _ := os.Getwd()
	type fields struct {
		fs    afero.Fs
		fsErr bool
	}
	type args struct {
		instance objects.Instance
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get Instance dir",
			args: args{instance: objects.Instance{Name: "test"}},
			want: dir + "/" + configAemInstancesDir + "/test/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			got, err := GetInstanceDirLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetInstanceDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetInstanceDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRunDirLocation(t *testing.T) {
	i := objects.Instance{Name: "1 2 3"}
	dir, _ := os.Getwd()
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
	}{
		{
			name:      "Get run location",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3/" + configAemRunDir + "/",
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get run location",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetRunDirLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetRunDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetRunDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPidFileLocation(t *testing.T) {
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
	}{
		{
			name:      "Get pid file location",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3" + "/" + configAemRunDir + "/" + configAemPidFile,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get pid file location error",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetPidFileLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetPidFileLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetPidFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAppDirLocation(t *testing.T) {
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
	}{
		{
			name:      "Get Instance app dir",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3" + "/" + configAemRunDir + "/" + configAppDir,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get Instance app dir",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetAppDirLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.GetAppDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.GetAppDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAemInstallDirLocation(t *testing.T) {
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
	}{
		{
			name:      "Get AEM install dir",
			args:      args{instance: i},
			want:      dir + `/` + configAemInstancesDir + "/" + "1-2-3/" + configAemRunDir + "/" + configAemInstallDir,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get AEM install dir error",
			args:      args{instance: i},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetAemInstallDirLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getAemInstallDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getAemInstallDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJarFileLocation(t *testing.T) {
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
	}{
		{
			name:      "Get jar location",
			want:      dir + "/" + configAemInstancesDir + "/" + configAemJar,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get jar location",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			jar := objects.AemJar{Version: "6.4"}
			got, err := GetJarFileLocation(&jar)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getJarFileLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getJarFileLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLogDirLocation(t *testing.T) {
	dir, _ := os.Getwd()

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
	}{
		{
			name:      "Log dir location",
			args:      args{instance: objects.Instance{Name: "123"}},
			want:      dir + "/" + configAemInstancesDir + "/123/" + configAemRunDir + "/" + configAemLogDir + "/",
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get log dir location error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetLogDirLocation(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogDirLocation error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLogDirLocation = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPackagesDirLocation(t *testing.T) {
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
	}{
		{
			name:      "Get pkg dir location",
			want:      dir + "/" + configAemInstancesDir + "/" + configPackageDir,
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get pkg dir location error",
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := getPackagesDirLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getPackagesDirLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getPackagesDirLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDirForPackage(t *testing.T) {
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
	}{
		{
			name:      "get dir for pkg",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0"}},
			want:      dir + "/" + configAemInstancesDir + "/" + configPackageDir + "/some-name/1.0.0",
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "get dir for pkg error",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0"}},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpTest()
			FilesystemError(tt.wantFsErr)
			got, err := GetDirForPackage(&tt.args.aemPackage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getDirForPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getDirForPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLocationForPackage(t *testing.T) {
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
	}{
		{
			name:      "Get pkg location",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0", DownloadName: "bla.zip"}},
			want:      dir + "/" + configAemInstancesDir + "/" + configPackageDir + "/some-name/1.0.0/bla.zip",
			wantErr:   false,
			wantFsErr: false,
		},
		{
			name:      "Get pkg location error",
			args:      args{aemPackage: objects.Package{Name: "some-name", Version: "1.0.0", DownloadName: "bla.zip"}},
			want:      ``,
			wantErr:   true,
			wantFsErr: true,
		},
	}
	for _, tt := range tests {
		setUpTest()
		t.Run(tt.name, func(t *testing.T) {
			FilesystemError(tt.wantFsErr)
			got, err := GetLocationForPackage(&tt.args.aemPackage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Structure.getLocationForPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Structure.getLocationForPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
