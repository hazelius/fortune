package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 8}, want: "6"},
		{name: "2", args: args{n: 100000}, want: "99634"},
		{name: "3", args: args{n: 999999999}, want: "999967331"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
