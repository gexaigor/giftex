package main

import (
	"log"

	"github.com/gexaigor/MyRestAPI/apiserver"
	_ "github.com/lib/pq"
)

func main() {
	config := apiserver.NewConfig()
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
