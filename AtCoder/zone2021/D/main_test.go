package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "ozRnonnoe"}, want: "zone"},
		{name: "2", args: args{s: "hellospaceRhellospace"}, want: ""},
		{name: "3", args: args{s: "abRadRbdRbd"}, want: "bd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
