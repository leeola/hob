package subproc

import (
	"bytes"
	"io"
	"net/http"
	"os/exec"

	"github.com/leeola/hob/middleware"
)

// Subproc is a convenience method for producing a synchronous process action.
func Subproc(bin string, args ...string) *SubprocAction {
	return NewSubprocAction(Config{
		Bin:  bin,
		Args: args,
	})
}

type Config struct {
	Bin      string
	Args     []string
	NoStdout bool
	NoStderr bool
}

type SubprocAction struct {
	config Config
}

func NewSubprocAction(c Config) *SubprocAction {
	return &SubprocAction{config: c}
}

func (a *SubprocAction) Act(w http.ResponseWriter, r *http.Request) {
	log := middleware.GetLog(r)
	cmd := exec.Command(a.config.Bin, a.config.Args...)

	var out bytes.Buffer
	if !a.config.NoStdout {
		cmd.Stdout = &out
	}
	if !a.config.NoStderr {
		cmd.Stderr = &out
	}

	// Don't immediately return the error, so we can write
	// whatever output we need to the event caller
	err := cmd.Run()
	if err != nil {
		log.Error("subproc action encountered error",
			"bin", a.config.Bin, "args", a.config.Args, "err", err.Error())

		// TODO(leeola): configure some way to extract process exit code into the http
		// exit code.
		//
		// Perhaps theres a small range of 500-599 codes we can use, so that, eg:
		// 	HTTP 531 == exit code 1
		// 	HTTP 532 == exit code 2
		// and so on.
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Make sure to write output, regardless of error.
	io.Copy(w, &out)
}
