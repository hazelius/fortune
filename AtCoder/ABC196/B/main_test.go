package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		x    string
		want string
	}{
		{name: "1", x: "5.90", want: "5"},
		{name: "2", x: "0", want: "0"},
		{name: "3", x: "84939825309432908832902189.9092309409809091329", want: "84939825309432908832902189"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.x); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
