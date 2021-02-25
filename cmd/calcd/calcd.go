package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli"
	"github.com/AnelBerik/Calculator2"
)

var (
	server *rpc.Server
)

func initServer(ctx *cli.Context) error {
	server = rpc.NewServer()
	return server.RegisterName(`calc`, &Calculator2.CalcService{})
}

func main() {
	app := cli.NewApp()
	app.Name = "calcd"
	app.Usage = "A JSON-RPC-based calculator service that supports addition and subtraction"

	app.Commands = []cli.Command{
		fifoCommand,
		tcpCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
