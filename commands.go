package hare

import (
	"errors"
)

func (server *Server) HandleCommand(cmd *Command) *Result {

	switch cmd.Name {
	default:
		return server.unknown()
	}

}

func (server *Server) unknown() *Result {
	return &Result{Err: errors.New("Unknown command"), Val: nil}
}
