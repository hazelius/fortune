package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	err := second()
	if errors.Cause(err) == strconv.ErrSyntax {
		fmt.Printf("%+v", err)
	} else {
		fmt.Printf("%v", err)
	}
}

func second() error {
	err := third()
	return err
}

func third() error {
	err := fourth()
	return err
}

func fourth() error {
	err := fifth("abc")
	if err != nil {
		return err
	}
	return nil
}

func fifth(val string) error {
	_, err := strconv.Atoi(val)
	if err != nil {
		// return errors.WithStack(strconv.ErrSyntax)
		return errors.Wrap(err, "Not Interger")
	}
	return nil
}
