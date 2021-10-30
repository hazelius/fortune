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
		want string
	}{
		{name: "1", args: args{r: strings.NewReader(`2 3
		1 2 3
		8 9 10`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`2 1
		1
		2`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`10 4
		1346 1347 1348 1349
		1353 1354 1355 1356
		1360 1361 1362 1363
		1367 1368 1369 1370
		1374 1375 1376 1377
		1381 1382 1383 1384
		1388 1389 1390 1391
		1395 1396 1397 1398
		1402 1403 1404 1405
		1409 1410 1411 1412`)}, want: "Yes"},
		{name: "4", args: args{r: strings.NewReader("1 1 0")}, want: "No"},
		{name: "5", args: args{r: strings.NewReader(`2 4
		13 14 15 16
		20 21 22 23`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
