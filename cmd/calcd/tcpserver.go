package main


import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli"
)


var tcpCommand = cli.Command{
	Name:    "tcp",
	Aliases: []string{"t"},
	Usage:   "Start the calculator service listening on the provided port, or 2000 by default",
	Action:  listenAndServeTCP,
	Before:  initServer,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "port,p",
			Value: 2000,
		},
	},
}


func listenAndServeTCP(c *cli.Context) error {

	port := c.Int(`port`)
	if port <= 0 {
		return fmt.Errorf("invalid port number, please specify a positive integer")
	}


	l, err := net.Listen(`tcp`, `:`+strconv.Itoa(port))
	if err != nil {
		return fmt.Errorf("Unable to listen for TCP connections on port %d: %v", port, err)
	}
	defer l.Close()
	log.Printf("Listening for connections on port %d...\n", port)


	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("Failed to accept incoming connection: %v", err)
		}
		codec := rpc.NewJSONCodec(conn)
		go server.ServeCodec(codec, 0)
	}
}
