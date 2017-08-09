package main

import (
	"io/ioutil"
	"log"
	"os"
	"golang.org/x/crypto/ssh"
)

func ssh_test() {
	ip := "example.com" //サーバのアドレス
	port := "22"            //ポート番号は文字列で指定
	user := "user"         //ユーザ名

	buf, err := ioutil.ReadFile("fileto/key")
	if err != nil {
		panic(err)
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Run("ls")
}
