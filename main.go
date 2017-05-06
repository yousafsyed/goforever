package main

import (
	"fmt"
	"github.com/takama/daemon"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	name        = "GoForever"
	description = "GoForever Service Command Deamonizer"
	port        = ":9977"
)

// dependencies that are NOT required by the service, but might be used
var dependencies = []string{"dummy.service"}

var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: GoForever install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			startCommand()
			return service.Start()
		case "stop":
			EndCommand()
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	for {
		select {
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			EndCommand()
			if killSignal == os.Interrupt {

				return "GoForever was interrupted by system signal", nil
			}

			return "GoForever was killed", nil
		}
	}
}

func init() {
	stdlog = log.New(os.Stdout, "", 0)
	errlog = log.New(os.Stderr, "", 0)

}

func startCommand() {

	if runtime.GOOS == "windows" {
		out, _ := exec.Command("cmd", "/C", config.Run.StartCommand).CombinedOutput()

		logOut(string(out))
	} else {
		out, _ := exec.Command("sh", "-c", config.Run.StartCommand).Output()

		logOut(string(out))
	}

}

func EndCommand() {

	if runtime.GOOS == "windows" {
		out, _ := exec.Command("cmd", "/C", config.Run.EndCommand).CombinedOutput()

		logOut(string(out))
	} else {
		out, _ := exec.Command("sh", "-c", config.Run.EndCommand).Output()

		logOut(string(out))
	}

}

func main() {
	srv, err := daemon.New(name, description, dependencies...)
	if err != nil {
		errlog.Println("Error: ", err)
		logOut(string(err.Error()))
		os.Exit(1)
	}
	startCommand()

	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		logOut(string(err.Error()))
		os.Exit(1)
	}
	fmt.Println(status)
	logOut(string(status))

}
