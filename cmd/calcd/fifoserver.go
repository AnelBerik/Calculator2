package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/AnelBerik/Calculator2"
	"github.com/urfave/cli"
)

var fifoCommand = cli.Command{
	Name:      "fifo",
	Aliases:   []string{"f"},
	Usage:     "Start the calculator service using named pipes / fifos as the transport",
	Action:    listenAndServeFifo,
	ArgsUsage: "[/path/to/request/pipe] [/path/to/request/pipe]",
	Before:    initServer,
}


func listenAndServeFifo(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("Invalid number of arguments.")
	}

	reqPath := c.Args().Get(0)
	respPath := c.Args().Get(1)

	for _, path := range []string{reqPath, respPath} {
		pathInfo, err := os.Stat(path)
		if err != nil {

			if err := syscall.Mkfifo(path, 0600); err != nil {
				return fmt.Errorf("Unable to create a named pipe at \"%s\": %v", path, err)
			}
		} else {

			if (pathInfo.Mode() & os.ModeNamedPipe) == 0 {
				return fmt.Errorf("A file at \"%s\" already exists and is not a named pipe", path)
			}
		}
	}

	conn := Calculator2.NewFifoConn(reqPath, respPath)
	defer conn.Close()
	log.Printf("Listening for requests at %s and sending responses to %s...\n", reqPath, respPath)


	for {
		codec := rpc.NewJSONCodec(conn)
		server.ServeCodec(codec, 0)
	}
}
