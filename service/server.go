package service

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Server struct {
	restartCmd string
	restartArgs []string
	statusCmd string
	statusArgs []string
}

func NewServer() *Server {
	srv := &Server{
		restartCmd: "systemctl",
		restartArgs: []string {"--user", "restart", "valheimserver.service"},
		statusCmd: "systemctl",
		statusArgs: []string {"--user", "status", "valheimserver.service"},
	}
	return srv
}

func (srv *Server) Restart() (string, error) {
	cmd := exec.Command(srv.restartCmd, srv.restartArgs[0:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	log.Println(out.String())
	return out.String(), nil
}

func (srv *Server) Status() (string, error) {
	cmd := exec.Command(srv.statusCmd, srv.statusArgs[0:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	log.Println(out.String())
	return out.String(), nil
}