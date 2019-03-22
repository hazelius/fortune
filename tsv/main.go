package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	in := `"first_name"	"last_name"	"username"
"Rob"	"Pike ""p"""	rob
# lines beginning with a # character are ignored
Ken	Thompson	"ken
and show"
"Robert\"	"Grie\\semer"	"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = '\t'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range records {
		for _, c := range line {
			fmt.Print(c + ", ")
		}
		fmt.Println(i)
	}
}
