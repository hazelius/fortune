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
		{name: "1", args: args{stdin: strings.NewReader(`1 1
		chokudai
		chokudai`)}, wantOut: `-1`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
		choku
		dai
		chokudai
		choku_dai`)}, wantOut: `choku__dai`},
		{name: "3", args: args{stdin: strings.NewReader(`2 2
		chokudai
		atcoder
		chokudai_atcoder
		atcoder_chokudai`)}, wantOut: `-1`},
		{name: "4", args: args{stdin: strings.NewReader(`4 4
		ab
		cd
		ef
		gh
		hoge
		fuga
		____
		_ab_cd_ef_gh_`)}, wantOut: `ab_cd_ef_gh`},
		{name: "5", args: args{stdin: strings.NewReader(`2 30
		b
		a
		a_b
		a__b
		a___b
		a____b
		a_____b
		a______b
		a_______b
		a________b
		a_________b
		a__________b
		a___________b
		a____________b
		a_____________b
		a______________b
		b_a
		b__a
		b___a
		b____a
		b_____a
		b______a
		b_______a
		b________a
		b_________a
		b__________a
		b___________a
		b____________a
		b_____________a`)}, wantOut: `b______________a`},
		{name: "6", args: args{stdin: strings.NewReader(`3 50
		b
		c
		a
		a_b_c
		a__b_c
		a___b_c
		a____b_c
		a_____b_c
		a______b_c
		a_______b_c
		a________b_c
		a_________b_c
		a__________b_c
		a___________b_c
		a____________b_c
		a_b__c
		a__b__c
		a___b__c
		a____b__c
		a_____b__c
		a______b__c
		a_______b__c
		a________b__c
		a_________b__c
		a__________b__c
		a___________b__c
		a____________b__c
		a_b___c
		a__b___c
		a___b___c
		a____b___c
		a_____b___c
		a______b___c
		a_______b___c
		a________b___c
		a_________b___c
		a__________b___c
		a___________b___c
		a____________b___c
		a_b____c
		a__b____c
		a___b____c
		a____b____c
		a_____b____c
		a______b____c
		a_______b____c
		a________b____c
		a_________b____c
		a__________b____c
		a___________b____c
		a____________b____c
		b__a`)}, wantOut: `a_b_____c`},
		{name: "7", args: args{stdin: strings.NewReader(`8 20
		a
		b
		c
		d
		e
		f
		g
		h
		a_b_c_d_e_f_g_h
		a_b_c_d_e_f_g__h
		a_b_c_d_e_f__g_h
		a_b_c_d_e__f_g_h
		a_b_c_d__e_f_g_h
		a_b_c__d_e_f_g_h
		a_b__c_d_e_f_g_h
		a__b_c_d_e_f_g_h
		a_b_c_d_e_f_h_g
		a_b_c_d_e_f_h__g
		a_b_c_d_e_f__h_g
		a_b_c_d_e__f_h_g
		a_b_c_d__e_f_h_g
		a_b_c__d_e_f_h_g
		a_b__c_d_e_f_h_g
		a__b_c_d_e_f_h_g`)}, wantOut: `a_b_c_d_e_g_f_h`},
		{name: "8", args: args{stdin: strings.NewReader(`1 1
		abcdefghijklmnol
		abcdefghijklmnol`)}, wantOut: `-1`},
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
