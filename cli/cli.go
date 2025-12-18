package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/hellotect2022go/nomadcoin/explorer"
	"github.com/hellotect2022go/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:		Set the PORT of the server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n\n")
	//os.Exit(0)
	runtime.Goexit() // defer 수행후 종료
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
