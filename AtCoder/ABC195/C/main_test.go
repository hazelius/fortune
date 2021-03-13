package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "1", n: 1010, want: "11"},
		{name: "2", n: 27182818284590, want: "107730272137364"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
