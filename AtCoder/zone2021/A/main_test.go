// HELLO SPACE
//

package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "abcdZONefghi"}, want: 1},
		{name: "2", args: args{s: "ZONeZONeZONe"}, want: 3},
		{name: "3", args: args{s: "helloAtZoner"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
