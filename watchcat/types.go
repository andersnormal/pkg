package watchcat

import (
	"net/http"

	"github.com/andersnormal/pkg/server"
)

// Opt ....
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Addr string
}

// Watchcat ...
type Watchcat interface {
	server.Listener
}

type watchcat struct {
	addr string

	http  *http.Server
	https *http.Server

	opts *Opts
}
