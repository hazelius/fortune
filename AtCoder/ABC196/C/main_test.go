package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "1", n: 33, want: 3},
		{name: "2", n: 1333, want: 13},
		{name: "3", n: 10000000, want: 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
