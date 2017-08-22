package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/yeka/zip"
)

func main() {
	var err error
	log.Printf("Start")

	flag.Parse()
	input_file := flag.Arg(0)
	output_file := flag.Arg(1)
	Read(input_file, output_file)

	Connect()
	conn, err := ssh_test()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	ftp_test(conn)

	err = Download("test_cydas", "test.txt")
	if err != nil {
		log.Printf("failed to download file, %v", err)
	}
	log.Printf("Finish !")
}

// Unzip zipファイルを解凍する
func Unzip(src, dest, password string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if password != "" {
			f.SetPassword(password)
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
