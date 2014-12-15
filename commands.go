package hare

import (
	"log"
)

func (server *Server) HandleCommand(cmd *Command) (res *Result) {
	switch cmd.Name {
	case "SEQ":
		res = server.seq(cmd)
	case "GET":
		res = server.get(cmd)
	case "SET":
		res = server.set(cmd)
	case "VGET":
		res = server.vget(cmd)
	case "VSET":
		res = server.vset(cmd)
	default:
		res = server.unknown()
	}
	return res
}

// { "name": "SEQ", "args": {} }
func (server *Server) seq(cmd *Command) *Result {
	return &Result{Val: server.db.Sequencer.Index}
}

// { "name": "GET", "args": { "id", id } }
func (server *Server) get(cmd *Command) *Result {
	var err error
	var res *Result = &Result{}

	var sId = cmd.Args["id"]
	if sId == "" {
		res.Err = "GET requires \"id\" arg"
		return res
	}

	var doc = &Document{}
	err = doc.SetKey(sId)
	if err != nil {
		res.Err = err.Error()
		return res
	}
	err = server.db.Load(doc)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	res.Val = doc
	return res
}

// { "name": "SET", "args": { "id": id, "body": body } }
func (server *Server) set(cmd *Command) *Result {
	var err error
	var res *Result = &Result{}
	var doc *Document = &Document{}

	var body = cmd.Args["body"]
	if body == "" {
		res.Err = "SET requires \"body\" arg"
		return res
	}

	err = doc.SetBody(body)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	var sId = cmd.Args["id"]
	if sId == "" {
		doc.K = server.db.Next()
		log.Println("no id provided, using", doc.K)
	} else {
		log.Println("id provided, using", sId)
		err = doc.SetKey(sId)
		if err != nil {
			res.Err = err.Error()
			return res
		}
		// TODO remove cache references to doc's id
	}

	log.Println("Saving doc", doc)
	err = server.db.Save(doc)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	res.Val = doc
	return res
}

// { "name": "VGET", "args": { "name": name } }
func (server *Server) vget(cmd *Command) *Result {
	var err error
	var res *Result = &Result{}

	var name = cmd.Args["name"]
	if name == "" {
		res.Err = "VGET requires \"name\" arg"
		return res
	}

	var view *View = &View{}
	err = view.SetKey(name)
	if err != nil {
		res.Err = err.Error()
		return res
	}
	err = server.db.Load(view)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	res.Val = view
	return res
}

// { "name": "VSET", "args": { "name": name, "src": src } }
func (server *Server) vset(cmd *Command) *Result {
	var err error
	var res *Result = &Result{}

	var name = cmd.Args["name"]
	if name == "" {
		res.Err = "VSET requires \"name\" arg"
		return res
	}

	var src = cmd.Args["src"]
	if src == "" {
		res.Err = "VSET requires \"src\" arg"
		return res
	}

	var view = &View{Src: src}
	err = view.SetKey(name)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	err = server.db.Save(view)
	if err != nil {
		res.Err = err.Error()
		return res
	}

	res.Val = view
	return res
}

func (server *Server) unknown() *Result {
	return &Result{Err: "Unknown command", Val: nil}
}
