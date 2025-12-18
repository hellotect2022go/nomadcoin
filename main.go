package main

import (
	"github.com/hellotect2022go/nomadcoin/cli"
	"github.com/hellotect2022go/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
