package main

import (
	log "github.com/Sirupsen/logrus"
	"proof/tasks"
	"time"
	"proof/plugs"
)

func main() {
	log.SetLevel(log.DebugLevel)

	pool := tasks.NewPool(nil)
	pool.Start()
	defer pool.Close()

	pl := &tasks.Plug{Name:"PLUGS"}
	pl.Register(3)
	pl.Schedule(100, time.Now(), &plugs.Payload{Name:"Independent", Value: "10"})

	select {}
}