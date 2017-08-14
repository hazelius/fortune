package main

import (
	"io/ioutil"
	"golang.org/x/crypto/ssh"
)

func ssh_test() (*ssh.Client, error) {
	ip := "devadmin.cydascloud.com" //サーバのアドレス
	port := "22"            //ポート番号は文字列で指定
	user := "ec2-user"         //ユーザ名

	buf, err := ioutil.ReadFile("C:\\Users\\yousu_000\\.ssh\\id_rsa.amazon")
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
