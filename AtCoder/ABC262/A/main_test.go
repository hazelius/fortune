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
		{name: "1", args: args{stdin: strings.NewReader(`2022`)}, wantOut: `2022`},
		{name: "2", args: args{stdin: strings.NewReader(`2023`)}, wantOut: `2026`},
		{name: "3", args: args{stdin: strings.NewReader(`3000`)}, wantOut: `3002`},
		{name: "4", args: args{stdin: strings.NewReader(`2024`)}, wantOut: `2026`},
		{name: "5", args: args{stdin: strings.NewReader(`2025`)}, wantOut: `2026`},
		{name: "6", args: args{stdin: strings.NewReader(`2026`)}, wantOut: `2026`},
		{name: "7", args: args{stdin: strings.NewReader(`2027`)}, wantOut: `2030`},
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
