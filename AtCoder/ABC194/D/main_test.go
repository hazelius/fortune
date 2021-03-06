package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "1", n: 2, want: "2"},
		{name: "2", n: 3, want: "4.5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
