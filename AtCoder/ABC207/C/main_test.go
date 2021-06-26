package main

import (
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{strings.NewReader(`3
		1 1 2
		2 2 3
		3 2 4`)}, want: 2},
		{name: "2", args: args{strings.NewReader(`19
		4 210068409 221208102
		4 16698200 910945203
		4 76268400 259148323
		4 370943597 566244098
		1 428897569 509621647
		4 250946752 823720939
		1 642505376 868415584
		2 619091266 868230936
		2 306543999 654038915
		4 486033777 715789416
		1 527225177 583184546
		2 885292456 900938599
		3 264004185 486613484
		2 345310564 818091848
		1 152544274 521564293
		4 13819154 555218434
		3 507364086 545932412
		4 797872271 935850549
		2 415488246 685203817`)}, want: 102},
		{name: "3", args: args{strings.NewReader(`3
		2 1 3
		3 2 3
		3 2 4`)}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
