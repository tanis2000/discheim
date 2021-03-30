package service

import (
	"log"
	"os/exec"
)

type Server struct {
	restartCmd string
	restartArgs []string
}

func NewServer() *Server {
	srv := &Server{
		//restartCmd: "ls",
		//restartArgs: []string {"-la"},
		restartCmd: "systemctl",
		restartArgs: []string {"restart valheimserver.service"},
	}
	return srv
}

func (srv *Server) Restart() error {
	cmd := exec.Command(srv.restartCmd, srv.restartArgs[0:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Print(string(out))
	return nil
}