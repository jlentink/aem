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
		wantReg string
	}{
		{
			name:    "Check if bin can be found",
			fields:  fields{bin: "vlt-lorum-ipsum"},
			wantErr: true,
		},
		{
			name:   "vlt response",
			fields: fields{bin: "vlt"},
			args: args{
				[]string{"--version"},
			},
			wantErr: false,
			wantReg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vlt{
				bin: tt.fields.bin,
			}
			_, err := v.execute(tt.args.params...)

			if (err != nil) != tt.wantErr {
				t.Errorf("vlt.execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			//if err == nil {
			//	//matched, err := regexp.MatchString(tt.wantReg, "seafood")
			//}
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
		// TODO: Add test cases.
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
