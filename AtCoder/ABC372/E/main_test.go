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
		{name: "1", args: args{stdin: strings.NewReader(`4 10
1 1 2
2 1 1
2 1 2
2 1 3
1 1 3
1 2 3
1 3 4
2 1 1
2 1 3
2 1 5`)}, wantOut: `2
1
-1
4
2
-1
`},
		{name: "2", args: args{stdin: strings.NewReader(`6 20
1 3 4
1 3 5
2 1 1
2 3 1
1 1 5
2 6 9
2 1 3
2 6 1
1 4 6
2 2 1
2 6 2
2 4 7
1 1 4
2 6 2
2 3 4
1 2 5
2 4 1
1 1 6
2 3 3
2 1 3`)}, wantOut: `1
5
-1
3
6
2
5
-1
5
3
6
4
4
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
