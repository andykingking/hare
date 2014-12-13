package hare

func (server *Server) HandleCommand(cmd *Command) *Result {

  return server.unknown()

}

func (server *Server) unknown() *Result {
  return &Result{Err: "Unknown command", Val: nil}
}
