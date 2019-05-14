package main

import (
	"reflect"
	"testing"
)

func Test_newVlt(t *testing.T) {
	tests := []struct {
		name string
		want vlt
	}{
		{
			name: "Get VLT",
			want: vlt{bin: "vlt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newVlt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newVlt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vlt_execute(t *testing.T) {
	type fields struct {
		bin string
	}
	type args struct {
		params []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantTxt string
	}{
		{
			name:    "Check if bin can be found",
			fields:  fields{bin: "vlt-lorum-ipsum"},
			wantErr: true,
		},
		{
			name:   "vlt response",
			fields: fields{bin: "echo"},
			args: args{
				[]string{"Testing VLT wrapper..."},
			},
			wantErr: false,
			wantTxt: "Testing VLT wrapper...\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vlt{
				bin: tt.fields.bin,
			}
			output, err := v.execute(tt.args.params...)

			if (err != nil) != tt.wantErr {
				t.Errorf("vlt.execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if tt.wantTxt != output {
					t.Errorf("vlt.execute() output = %v, wantTxt %v", output, tt.wantTxt)
				}
			}
		})
	}
}

func Test_vlt_checkBin(t *testing.T) {
	type fields struct {
		bin string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Try to find not existing file",
			fields: fields{bin: "vlt-3323212"},
			want:   false,
		},
		{
			name:   "Try to find existing file",
			fields: fields{bin: "ls"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vlt{
				bin: tt.fields.bin,
			}
			if got := v.checkBin(); got != tt.want {
				t.Errorf("vlt.checkBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vlt_getPaths(t *testing.T) {
	type fields struct {
		bin string
	}
	type args struct {
		paramPath   string
		configPaths []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []vltPath
	}{
		{
			name: "Test param path",
			args: args{paramPath: "/content", configPaths: []string{}},
			want: []vltPath{{path: "/content", recursive: false}},
		},
		{
			name: "Test param path",
			args: args{paramPath: "", configPaths: []string{"/foo", "/bar"}},
			want: []vltPath{{path: "/foo", recursive: false}, {path: "/bar", recursive: false}},
		},
		{
			name: "Test param path",
			args: args{paramPath: "/content", configPaths: []string{"/foo", "/bar"}},
			want: []vltPath{{path: "/content", recursive: false}},
		},
		{
			name: "Test param path",
			args: args{paramPath: "!/content", configPaths: []string{"/foo", "/bar"}},
			want: []vltPath{{path: "/content", recursive: true}},
		},
		{
			name: "Test param path",
			args: args{paramPath: "", configPaths: []string{"!/foo", "/bar"}},
			want: []vltPath{{path: "/foo", recursive: true}, {path: "/bar", recursive: false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vlt{
				bin: tt.fields.bin,
			}
			if got := v.getPaths(tt.args.paramPath, tt.args.configPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vlt.getPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vlt_pathDetail(t *testing.T) {
	type fields struct {
		bin string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   vltPath
	}{
		{
			name: "Check non-recursive path",
			args: args{path: "/content"},
			want: vltPath{path: "/content", recursive: false},
		},
		{
			name: "Check recursive path",
			args: args{path: "!/content"},
			want: vltPath{path: "/content", recursive: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vlt{
				bin: tt.fields.bin,
			}
			if got := v.pathDetail(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vlt.pathDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
