package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func Read(file_name string, output_file string) {

	//読み込みファイル準備
	file1, err := os.Open(file_name)
	failOnError(err)
	defer file1.Close()

	//書き込みファイル準備
	file2, err := os.Create(output_file)
	failOnError(err)
	defer file2.Close()

	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))
	//  reader := csv.NewReader(transform.NewReader(file1, japanese.EUCJP.NewDecoder()))
	//  reader := csv.NewReader(file1) //utf8
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	writer := csv.NewWriter(transform.NewWriter(file2, japanese.ShiftJIS.NewEncoder()))
	//  writer := csv.NewWriter(transform.NewWriter(file2, japanese.EUCJP.NewEncoder()))
	//  writer := csv.NewWriter(file2) //utf8
	writer.UseCRLF = true //デフォルトはLFのみ
	//  writer.Comma = '\t'

	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}
		var new_record []string
		for i, v := range record {
			if i > 0 {
				new_record = append(new_record, fmt.Sprint(i)+":"+v)
			}
		}
		writer.Write(new_record) // 1行書き出す
		//      log.Printf("%#v", record[0] + "," + record[1])
	}
	writer.Flush()
}
