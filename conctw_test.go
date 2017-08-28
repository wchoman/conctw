package conctw

import (
	"testing"

	"golang.org/x/tour/tree"
)

func Test_recurseWalk(t *testing.T) {
	type args struct {
		t  *tree.Tree
		ch chan int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recurseWalk(tt.args.t, tt.args.ch)
		})
	}
}

func TestWalk(t *testing.T) {
	type args struct {
		t  *tree.Tree
		ch chan int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Walk(tt.args.t, tt.args.ch)
		})
	}
}

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"Simple",
			args{
				tree.New(1),
				tree.New(1),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
