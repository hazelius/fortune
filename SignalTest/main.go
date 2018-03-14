package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exitChan <- 0

			case os.Interrupt:
				fmt.Println("Interrupt")
				exitChan <- 0

			default:
				fmt.Println("Unknown signal.")
				exitChan <- 1
			}
		}
	}()

	go func() {
		for {
			fmt.Println("start...")
			time.Sleep(5 * time.Second)
			fmt.Println("end...")
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

// func main() {
// 	// サイズが1より大きいチャネルを作成
// 	signals := make(chan os.Signal, 1)
// 	// 最初のチャネル以降は、可変長引数で任意の数のシグナルを設定可能
// 	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
// 	s := <-signals
// 	switch s {
// 	case syscall.SIGINT:
// 		fmt.Println("SIGINT")
// 	case syscall.SIGTERM:
// 		fmt.Println("SIGTERM")
// 	}
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("runnning")
// }
