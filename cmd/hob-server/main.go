package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/leeola/hob"
	"github.com/leeola/hob/actions/subproc"
	"github.com/leeola/hob/server"
)

func main() {
	// parse the config path option
	var configPath string
	flag.StringVar(&configPath, "config", "./config.toml", "path to hob config")
	flag.Parse()

	f, err := os.Open(configPath)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var conf Config
	if _, err := toml.DecodeReader(f, &conf); err != nil {
		panic(err)
	}

	actions := map[string]hob.Action{}

	for _, a := range conf.Actions.Subprocs {
		if _, ok := actions[a.Action]; ok {
			panic("two actions share the same name: " + a.Action)
		}
		actions[a.Action] = subproc.Subproc(a.Bin, a.Args...)
	}

	server.ListenAndServe(conf.BindAddr, hob.Config{
		Events:  conf.Events,
		Actions: actions,
	})
}
