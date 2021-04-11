package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: "1210"}, want: "Yes"},
		{name: "2", args: args{n: "777"}, want: "Yes"},
		{name: "3", args: args{n: "123456789"}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
