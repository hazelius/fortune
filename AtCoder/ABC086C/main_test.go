package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		txy [][3]int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{txy: [][3]int64{{3, 1, 2}, {6, 1, 1}}}, want: "Yes"},
		{name: "2", args: args{txy: [][3]int64{{2, 100, 100}}}, want: "No"},
		{name: "3", args: args{txy: [][3]int64{{5, 1, 1}, {100, 1, 1}}}, want: "No"},
		{name: "4", args: args{txy: [][3]int64{{2, 1, 1}, {7, 4, 1}}}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.txy); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
