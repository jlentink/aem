package main

import (
	"testing"
)

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

func Test_sliceUtil_sliceCompareString(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		s    *sliceUtil
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
			s := &sliceUtil{}
			if got := s.sliceStringCompare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("sliceUtil.sliceCompareString() = %v, want %v", got, tt.want)
			}
		})
	}
}
