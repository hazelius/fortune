package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		y string
		m int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{y: "22", m: 10}, want: "2"},
		{name: "2", args: args{y: "999", m: 1500}, want: "3"},
		{name: "3", args: args{y: "100000000000000000000000000000000000000000000000000000000000", m: 1000000000000000000}, want: "1"},
		{name: "4", args: args{y: "11", m: 99000000}, want: "98999998"},
		{name: "5", args: args{y: "11", m: 999999999999999999}, want: "999999999999999997"},
		{name: "6", args: args{y: "10", m: 1000000000000000000}, want: "999999999999999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.y, tt.args.m); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		x   string
		m   int
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{x: "22", m: 10, min: 3, max: 10}, want: 4},
		{name: "2", args: args{x: "999", m: 1500, min: 10, max: 1500}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.x, tt.args.m, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatInt(t *testing.T) {
	type args struct {
		i    int
		base int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{i: 111, base: 2}, want: []int{1, 1, 0, 1, 1, 1, 1}},
		{name: "2", args: args{i: 27, base: 3}, want: []int{1, 0, 0, 0}},
		{name: "3", args: args{i: 30, base: 3}, want: []int{1, 0, 1, 0}},
		{name: "4", args: args{i: 9, base: 4500}, want: []int{9}},
		// {name: "5", args: args{i: 1500, base: 1401}, want: "9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatInt(tt.args.i, tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
