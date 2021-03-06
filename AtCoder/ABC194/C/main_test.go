package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: []int{2, 8, 4}}, want: "56"},
		{name: "2", args: args{r: []int{-5, 8, 9, -4, -3}}, want: "950"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
