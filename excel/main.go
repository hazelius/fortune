package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()

	if err := f.SaveAs("サンプル.xlsx"); err != nil {
		fmt.Println(err)
	}
}
