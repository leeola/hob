package main

import (
	"github.com/leeola/hob"
	"github.com/leeola/hob/actions/subproc"
	"github.com/leeola/hob/server"
)

func main() {
	server.ListenAndServe(":4001", hob.Config{
		Events: map[string]string{
			"hello": "say hello",
			"fail":  "exit 7",
		},
		Actions: map[string]hob.Action{
			"say hello": subproc.Subproc("echo", "hello sir or madam"),
			"exit 7":    subproc.Subproc("bash", "-c", "exit 7"),
		},
	})
}
