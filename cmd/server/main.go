package main

import (
	"fmt"

	"github.com/sergioc0sta/hammer/internal/cli"
)

func main() {

	prs := cli.InsertParameters()

	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrencies:", prs.Concurrencies)

	fmt.Println("Server is starting...")
}
