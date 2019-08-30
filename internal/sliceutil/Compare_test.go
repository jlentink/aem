package sliceutil

import "testing"

func TestSliceStringCompare(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal slices",
			args: args{
				s1: []string{"1", "2"},
				s2: []string{"1", "2"},
			},
			want: true,
		},
		{
			name: "different length long 1",
			args: args{
				s1: []string{"1", "2"},
				s2: []string{"1", "2", "3"},
			},
			want: false,
		},
		{
			name: "different length long 2",
			args: args{
				s1: []string{"1", "2", "3"},
				s2: []string{"1", "2"},
			},
			want: false,
		},
		{
			name: "Different value",
			args: args{
				s1: []string{"1", "2", "3"},
				s2: []string{"1", "2", "4"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCompare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("StringCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
