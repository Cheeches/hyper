package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/cognicraft/hyper"
)

func main() {
	q := flag.String("q", ".", "Query")
	flag.Parse()

	c := hyper.NewClient()
	item, err := c.Fetch(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	res := hyper.Query(item, *q)
	bs, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
