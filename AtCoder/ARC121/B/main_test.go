package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		b int64
		c int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{b: 11, c: 2}, want: "3"},
		{name: "2", args: args{b: 0, c: 4}, want: "4"},
		{name: "3", args: args{b: 112, c: 20210213}, want: "20210436"},
		{name: "4", args: args{b: -211, c: 1000000000000000000}, want: "1000000000000000422"},
		{name: "5", args: args{b: 1000000000000000000, c: 1}, want: "2"},
		{name: "6", args: args{b: 0, c: 1}, want: "1"},
		{name: "7", args: args{b: 3, c: 14}, want: "19"},
		{name: "8", args: args{b: 489647926824927166, c: 979295853649854331}, want: "1958591707299708661"},
		{name: "9", args: args{b: 489647926824927166, c: 979295853649854332}, want: "1958591707299708663"},
		{name: "10", args: args{b: 489647926824927166, c: 979295853649854333}, want: "1958591707299708664"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
