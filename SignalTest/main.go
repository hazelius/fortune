package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	exitCh := make(chan struct{})
	go func(ctx context.Context) {
		for {
			fmt.Println("start...")
			time.Sleep(3 * time.Second)
			fmt.Println("end...")

			select {
			case <-ctx.Done():
				fmt.Println("received done, exiting in 500 milliseconds")
				time.Sleep(500 * time.Millisecond)
				exitCh <- struct{}{}
				return
			default:
			}
		}
	}(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	go func() {
		select {
		case s := <-signalCh:
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
			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
			case os.Interrupt:
				fmt.Println("Interrupt")
			default:
				fmt.Println("Unknown signal.")
			}
			cancel()
			return
		}
	}()
	<-exitCh
}
