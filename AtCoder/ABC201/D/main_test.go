package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		h  int
		w  int
		as []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{h: 3, w: 3, as: []string{
			"---",
			"+-+",
			"+--",
		}}, want: "Takahashi"},
		{name: "2", args: args{h: 2, w: 4, as: []string{
			"+++-",
			"-+-+",
		}}, want: "Aoki"},
		{name: "3", args: args{h: 1, w: 1, as: []string{"-"}}, want: "Draw"},
		{name: "4", args: args{h: 2, w: 2, as: []string{
			"++",
			"++",
		}}, want: "Draw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.h, tt.args.w, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
