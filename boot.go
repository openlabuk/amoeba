package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Startup() {
	fmt.Println("Amoeba started.")
	fmt.Println("listing on:", server+":"+fmt.Sprint(port))
}

func GracefulShutdown() {
	s := make(chan os.Signal, 1)

	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)

	go func() {
		<-s
		fmt.Println("Shutting down...")
		os.Exit(0)
	}()
}
