package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		ms []mountain
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 3, ms: []mountain{{"Everest", 8849}, {"K2", 8611}, {"Kangchenjunga", 8586}}}, want: "K2"},
		{name: "2", args: args{n: 4, ms: []mountain{{"Kita", 3193}, {"Aino", 3189}, {"Fuji", 3776}, {"Okuhotaka", 3190}}}, want: "Kita"},
		{name: "3", args: args{n: 4, ms: []mountain{
			{"QCFium", 2846},
			{"chokudai", 2992},
			{"kyoprofriends", 2432},
			{"penguinman", 2390},
		}}, want: "QCFium"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.ms); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
