package main

import (
	"context"
	"fmt"

	"github.com/andersnormal/pkg/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := server.NewServer(ctx)

	d := server.NewDebugListener(
		server.WithPprof(),
		server.WithStatusAddr(":8443"),
	)
	s.Listen(d)

	if err := s.Wait(); err != nil {
		fmt.Println(err)
	}
}
