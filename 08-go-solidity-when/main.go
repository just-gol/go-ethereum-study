package main

import (
	"log"

	"08-go-solidity-when/bootstrap"
	"08-go-solidity-when/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r, err := bootstrap.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = r.Run()
}
