package main

import (
	"fmt"

	"github.com/sergioc0sta/hammer/internal/cli"
	"github.com/sergioc0sta/hammer/internal/infra/handler"
)

func main() {

	prs, err := cli.InsertParameters()

	if(err != nil){
    panic(err)
	}

	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrency:", prs.Concurrency)
	status, err := handler.ServerRequestHandler(prs.URL)

	fmt.Println("STATUS:", status)
	fmt.Println("ERR:", err)

	fmt.Println("Server is starting...")
}
