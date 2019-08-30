package aem

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"reflect"
	"testing"
)

func TestGetByName(t *testing.T) {
	instances := []objects.Instance{
		{
			Name: "Some-name",
		},
		{
			Name: "Some-name #2",
		},
		{
			Name: "Some-name #3",
		},
	}

	type args struct {
		n string
		i []objects.Instance
	}
	tests := []struct {
		name    string
		args    args
		want    *objects.Instance
		wantErr bool
	}{
		{
			name:    "Find instance #1",
			args:    args{n: "Some-name #2", i: instances},
			want:    &objects.Instance{Name: "Some-name #2"},
			wantErr: false,
		},
		{
			name:    "Find instance #2",
			args:    args{n: "Some-names", i: instances},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByName(tt.args.n, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetByGroup(t *testing.T) {
	instances := []objects.Instance{
		{
			Name:  "Some-name",
			Group: "Local",
		},
		{
			Name:  "Some-name #2",
			Group: "development",
		},
		{
			Name:  "Some-name #3",
			Group: "development",
		},
	}
	type args struct {
		g string
		i []objects.Instance
	}
	tests := []struct {
		name string
		args args
		want []objects.Instance
	}{
		{
			name: "Find one in group",
			args: args{g: "Local", i: instances},
			want: []objects.Instance{{Name: "Some-name", Group: "Local"}},
		},
		{
			name: "Find two in group",
			args: args{g: "development", i: instances},
			want: []objects.Instance{
				{Name: "Some-name #2", Group: "development"},
				{Name: "Some-name #3", Group: "development"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got, err := GetByGroup(tt.args.g, tt.args.i); !reflect.DeepEqual(got, tt.want)
			//if got {
			//	t.Errorf("GetByGroup() = %v, want %v", got, tt.want)
			//}
		})
	}
}
