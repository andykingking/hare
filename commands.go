package hare

import (
	"errors"
)

func (server *Server) HandleCommand(cmd *Command) (res *Result) {
	switch cmd.Name {
	default:
		res = server.unknown()
	}
	return res
}

}

func (server *Server) unknown() *Result {
	return &Result{Err: errors.New("Unknown command"), Val: nil}
}
