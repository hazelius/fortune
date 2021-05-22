package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "16", args: args{a: 21, b: 30, k: 1}, want: "aaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
		// {name: "17", args: args{a: 4, b: 4, k: 38}, want: "abaababb"},
		// {name: "18", args: args{a: 4, b: 4, k: 39}, want: "abaabbab"},
		// {name: "19", args: args{a: 4, b: 4, k: 40}, want: "abaabbba"},
		// {name: "19", args: args{a: 4, b: 4, k: 41}, want: "abaabbba"},
		// {name: "19", args: args{a: 4, b: 4, k: 42}, want: "abaabbba"},
		// {name: "19", args: args{a: 4, b: 4, k: 43}, want: "abaabbba"},
		// {name: "1", args: args{a: 2, b: 2, k: 4}, want: "baab"},
		{name: "2", args: args{a: 30, b: 30, k: 118264581564861424}, want: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{name: "3", args: args{a: 30, b: 30, k: 1}, want: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
	}
	// for i := 11811564861404; i < 11811564861424; i++ {
	// 	fmt.Println(run(21, 30, i))
	// }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
