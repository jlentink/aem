package output

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Loglevel Normal",
			args: args{level: NORMAL},
		},
		{
			name: "Test Loglevel Verbose",
			args: args{level: VERBOSE},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.level)
			if logLevel != tt.args.level {
				t.Errorf("SetLevel() error = %v, want %v", logLevel, tt.args.level)
				return
			}
		})
	}
}

func TestSetVerbose(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Set verbose level=true",
			args: args{v: true},
			want: VERBOSE,
		},
		{
			name: "Set verbose level=false",
			args: args{v: false},
			want: NORMAL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(NORMAL)
			SetVerbose(tt.args.v)
			if tt.want != logLevel {
				t.Errorf("SetVerbose() error = %b, want %b", logLevel, tt.want)
				return
			}
		})
	}
}

func TestPrint(t *testing.T) {
	type args struct {
		l int
		o string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantErr   bool
		wantLevel int
	}{
		{
			name:      "Test normal print",
			args:      args{l: NORMAL, o: "a fox is a wolf who sends flowers\n"},
			want:      34,
			wantErr:   false,
			wantLevel: NORMAL,
		},
		{
			name:      "Test verbose print on normal level",
			args:      args{l: VERBOSE, o: "a fox is a wolf who sends flowers\n"},
			want:      0,
			wantErr:   false,
			wantLevel: NORMAL,
		},
		{
			name:      "Test verbose print on verbose level",
			args:      args{l: VERBOSE, o: "a fox is a wolf who sends flowers\n"},
			want:      34,
			wantErr:   false,
			wantLevel: VERBOSE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.wantLevel)
			got, err := Print(tt.args.l, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Print() output = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	type args struct {
		l int
		o string
		a []string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantErr   bool
		wantLevel int
	}{
		{
			name:      "Test normal print",
			args:      args{l: NORMAL, o: "a fox is a wolf who sends %s\n", a: []string{"flowers"}},
			want:      36,
			wantErr:   false,
			wantLevel: NORMAL,
		},
		{
			name:      "Test verbose print on normal level",
			args:      args{l: VERBOSE, o: "a fox is a wolf who sends %s\n", a: []string{"flowers"}},
			want:      0,
			wantErr:   false,
			wantLevel: NORMAL,
		},
		{
			name:      "Test verbose print on verbose level",
			args:      args{l: VERBOSE, o: "a fox is a wolf who sends %s\n", a: []string{"flowers"}},
			want:      36,
			wantErr:   false,
			wantLevel: VERBOSE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.wantLevel)
			got, err := Printf(tt.args.l, tt.args.o, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("Printf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Printf() = %v, want %v", got, tt.want)
			}
		})
	}
}
