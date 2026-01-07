package main

import (
	"fmt"

	"github.com/sergioc0sta/hammer/internal/cli"
)

func main() {

	prs, err := cli.InsertParameters()

	if(err != nil){
    panic(err)
	}

	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrency:", prs.Concurrency)

	fmt.Println("Server is starting...")
}
