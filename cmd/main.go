package main

import (
	"fmt"
	"os"

	"github.com/Fheyrian/test-api/httpsrv"
)

func main() {
	srv := httpsrv.Server{
		ListenIP:   "127.0.0.1",
		ListenPort: "8080",
		DBconnStr:  "mysql://root:password@127.0.0.1:3306/test",
	}

	fmt.Println("Starting http server...")
	if err := srv.Init(); err != nil {
		fmt.Println("error initializing server")
		fmt.Println(err)
		os.Exit(1)
	}
	if err := srv.Start(); err != nil {
		fmt.Println(err)
	}
}
