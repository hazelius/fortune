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
		{name: "1", args: args{r: strings.NewReader(`2
		4 0 1
		5 3
		2`)}, want: 6},
		{name: "2", args: args{r: strings.NewReader(`1
		5`)}, want: 5},
		{name: "3", args: args{r: strings.NewReader(`5
		900606388 317329110 665451442 1045743214 260775845 726039763 57365372 741277060 944347467
		369646735 642395945 599952146 86221147 523579390 591944369 911198494 695097136
		138172503 571268336 111747377 595746631 934427285 840101927 757856472
		655483844 580613112 445614713 607825444 252585196 725229185
		827291247 105489451 58628521 1032791417 152042357
		919691140 703307785 100772330 370415195
		666350287 691977663 987658020
		1039679956 218233643
		70938785`)}, want: 1073289207},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
