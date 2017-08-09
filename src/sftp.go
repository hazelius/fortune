package main

import (
  "log"
  "golang.org/x/crypto/ssh"
	"github.com/pkg/sftp"
)

func ftp_test(conn *ssh.Client) {

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()

	// walk a directory
	w := sftp.Walk("/home/user")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	// leave your mark
	f, err := sftp.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}

	// check it's there
	fi, err := sftp.Lstat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)
}
