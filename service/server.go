package service

import (
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
		//restartCmd: "ls",
		//restartArgs: []string {"-la"},
		restartCmd: "systemctl",
		restartArgs: []string {"--user restart valheimserver.service"},
		statusCmd: "systemctl",
		statusArgs: []string {"--user status valheimserver.service"},
	}
	return srv
}

func (srv *Server) Restart() (string, error) {
	cmd := exec.Command(srv.restartCmd, srv.restartArgs[0:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Print(string(out))
	return string(out), nil
}

func (srv *Server) Status() (string, error) {
	cmd := exec.Command(srv.statusCmd, srv.statusArgs[0:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Print(string(out))
	return string(out), nil
}