package main

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"
)

func TestSecond(t *testing.T) {
	err := second()
	if errors.Cause(err) == strconv.ErrSyntax {
		t.Error("エラーが発生していない")
	}
}

func TestThird(t *testing.T) {
	err := third()
	if errors.Cause(err) == strconv.ErrSyntax {
		t.Error("エラーが発生していない")
	}
}

func TestFourth(t *testing.T) {
	err := fourth()
	if errors.Cause(err) == strconv.ErrSyntax {
		t.Error("エラーが発生していない")
	}
}

func TestFifth(t *testing.T) {
	err := fifth("aaa")
	if errors.Cause(err) == strconv.ErrSyntax {
		t.Error("エラーが発生していない")
	}

	err = fifth("123")
	if err != nil {
		t.Errorf("数値が変換できていない %v", err)
	}
}
