package main

import (
	"assessment/presentation/grpc"
	"os"
	"syscall"

	"os/signal"

	"assessment/pkg"
)

func main() {
	pkg.LoadConfig()
	pkg.InitLog()

	grpc.Start()

	// wait for `Ctrl+c` or docker stop/restart signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch

	// Stop the application
	grpc.Stop()
}
