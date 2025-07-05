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
		{name: "1", args: args{stdin: strings.NewReader(`5
1 2 3
1 4 5
2 3
1 6 2
2 5`)}, wantOut: `11
19
`},
		{name: "2", args: args{stdin: strings.NewReader(`10
1 75 22
1 81 72
1 2 97
1 84 82
1 2 32
1 39 57
2 45
1 40 16
2 32
2 42`)}, wantOut: `990
804
3024
`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 160449218 954291757
2 17217760
1 353195922 501899080
1 350034067 910748511
1 824284691 470338674
2 180999835
1 131381221 677959980
1 346948152 208032501
1 893229302 506147731
2 298309896`)}, wantOut: `16430766442004320
155640513381884866
149721462357295680
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
