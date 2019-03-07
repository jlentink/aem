package main

import (
	"github.com/kataras/iris/core/errors"
	"testing"
)

func Test_exitProgram(t *testing.T) {
	type args struct {
		message string
		args    []interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "Exit program",
			args: args{
				message: "Stop now",
				args:    []interface{}{1},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("exitProgram() paniced when wantPanic = %v", tt.wantPanic)
				}
			}()
			exitProgram(tt.args.message, tt.args.args...)
		})
	}
}

func Test_exitFatal(t *testing.T) {
	type args struct {
		err     error
		message string
		args    []interface{}
	}
	tests := []struct {
		name string
		args args
		wantPanic bool
		want string
	}{
		{
			name: "No Err",
			args: args{
				err: nil,
				message: "no Error",
				args: []interface{}{1, 2, 3},
			},
			wantPanic: false,
		},
		{
			name: "Err",
			args: args{
				err: errors.New("Some Error"),
				message: "Error",
				args: []interface{}{1, 2, 3},
			},
			wantPanic: true,
		},
		{
			name: "Err",
			args: args{
				err: errors.New("Some Error"),
				message: "Error",
				args: []interface{}{},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("exitFatal() paniced when wantPanic = %v", tt.wantPanic)
				}
			}()
			exitFatal(tt.args.err, tt.args.message, tt.args.args...)
		})
	}
}
