package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		r [][]int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: [][]int64{{3, 9, 5}, {4, 8, 5}, {5, 7, 5}}}, want: "8"},
		{name: "2", args: args{r: [][]int64{{5, 9, 5}, {6, 8, 5}, {7, 5, 5}}}, want: "-1"},
		{name: "3", args: args{r: [][]int64{
			{158260522, 877914575, 602436426},
			{24979445, 861648772, 623690081},
			{433933447, 476190629, 262703497},
			{211047202, 971407775, 628894325},
			{731963982, 822804784, 450968417},
			{430302156, 982631932, 161735902},
			{880895728, 923078537, 707723857},
			{189330739, 910286918, 802329211},
			{404539679, 303238506, 317063340},
			{492686568, 773361868, 125660016},
		}}, want: "861648772"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
