package main

import "testing"

func Test_run(t *testing.T) {
	var maxim string

	for i := 1; i < 100000; i++ {
		if i%2 == 0 {
			maxim += "8"
		} else {
			maxim += "0"
		}
	}

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "0601889"}, want: "6881090"},
		{name: "2", args: args{s: "86910"}, want: "01698"},
		{name: "3", args: args{s: "01010"}, want: "01010"},
		{name: "4", args: args{s: "1000611199998888"}, want: "8888666611190001"},
		{name: "5", args: args{s: "6661010100"}, want: "0010101999"},
		{name: "6", args: args{s: maxim}, want: maxim},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
