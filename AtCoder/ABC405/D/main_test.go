package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		stdin io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{name: "1", args: args{stdin: strings.NewReader(`3 4
...E
.#..
....`)}, wantOut: `>>>E
^#>^
>>>^
`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
##
##
##`)}, wantOut: `##
##
##
`},
		{name: "3", args: args{stdin: strings.NewReader(`7 20
....................
..#..#..####..#E##..
..#..#..#..#..#.....
..E###..#..#..####..
.....#..#..E.....#..
.....#..####..####..
....................`)}, wantOut: `>v<<<<<>>>>>>>>v<<<<
>v#^<#^^####v^#E##^<
>v#^<#v^#>v#v<#^<<<<
>>E###v<#>v#v<####^<
>>^<<#v<#>>E<<<<<#^<
>>^<<#v<####^<####^<
>>^<<<<<>>>>^<<<<<^<
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.stdin, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
