package main

import "testing"

func Test_cased(t *testing.T) {
	type args struct {
		l int64
		r int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{l: 2, r: 6}, want: 6},
		{name: "2", args: args{l: 0, r: 0}, want: 1},
		{name: "3", args: args{l: 1000000, r: 1000000}, want: 0},
		{name: "4", args: args{l: 12345, r: 67890}, want: 933184801},
		{name: "4", args: args{l: 0, r: 1000000}, want: 500001500001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cased(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("cased() = %v, want %v", got, tt.want)
			}
		})
	}
}
