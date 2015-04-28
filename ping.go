package main

import (
	"time"

	"./common"
	"./logs"
)

func Ping() {
	ticker := time.Tick(config.Docker.Health * time.Second)
	for _ = range ticker {
		if err := common.Docker.Ping(); err != nil {
			logs.Fatal(err)
		}
	}
}