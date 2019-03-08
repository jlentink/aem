package main

import "testing"

func TestSliceUtil_InSliceInt64(t *testing.T) {
	type args struct {
		slice  []int64
		needle int64
	}
	tests := []struct {
		name string
		s    *sliceUtil
		args args
		want bool
	}{
		{
			name: "Find 1 in slice",
			args: args{
				slice:  []int64{1, 2, 3, 4, 5, 6, 7, 8},
				needle: 1,
			},
			want: true,
		},
		{
			name: "Find 11 in slice",
			args: args{
				slice:  []int64{1, 2, 3, 4, 5, 6, 7, 8},
				needle: 11,
			},
			want: false,
		},
		{
			name: "Find 9 in slice",
			args: args{
				slice:  []int64{1, 2, 3, 4, 5, 6, 7, 8},
				needle: 9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sliceUtil{}
			if got := s.inSliceInt64(tt.args.slice, tt.args.needle); got != tt.want {
				t.Errorf("sliceUtil.inSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceUtil_InSliceString(t *testing.T) {
	type args struct {
		slice  []string
		needle string
	}
	tests := []struct {
		name string
		s    *sliceUtil
		args args
		want bool
	}{
		{
			name: "Find a in slice",
			args: args{
				slice:  []string{"a", "aa", "bbb", "ccc"},
				needle: "a",
			},
			want: true,
		},
		{
			name: "Find aa in slice",
			args: args{
				slice:  []string{"a", "aa", "bbb", "ccc"},
				needle: "aa",
			},
			want: true,
		},
		{
			name: "Find aaa in slice",
			args: args{
				slice:  []string{"a", "aa", "bbb", "ccc"},
				needle: "aaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sliceUtil{}
			if got := s.inSliceString(tt.args.slice, tt.args.needle); got != tt.want {
				t.Errorf("sliceUtil.inSliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
