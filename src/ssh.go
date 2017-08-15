package main

import (
	"io/ioutil"
	"golang.org/x/crypto/ssh"
)

func ssh_test() (*ssh.Client, error) {
	ip := "example.com" //サーバのアドレス
	port := "22"            //ポート番号は文字列で指定
	user := "user"         //ユーザ名

	buf, err := ioutil.ReadFile("path/to/key")
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
	return conn, err

	// defer conn.Close()
	//
	// session, err := conn.NewSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer session.Close()
	//
	// out, err := session.CombinedOutput("ls")
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(string(out))
}
