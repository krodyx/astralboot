package main

// A channel muxer for the event system

import (
	"math/rand"
	"time"
	//	"fmt"
	"github.com/dustin/go-broadcast"
)

type Events struct {
	caster broadcast.Broadcaster
}

type notif struct {
	Name   string
	Status string
}

func NewEvents() (e *Events) {
	e = &Events{}
	e.caster = broadcast.NewBroadcaster(1024)
	// random testing
	//go e.InsertRandom()
	return
}

func (e *Events) GetListener() chan interface{} {
	logger.Info("Create Listener")
	listener := make(chan interface{})
	e.caster.Register(listener)
	return listener
}

func (e *Events) CloseListener(listener chan interface{}) {
	logger.Info("Closing Listener")
	e.caster.Unregister(listener)
	close(listener)
}

func (e *Events) Insert(section string, status string) {
	n := &notif{
		Name:   section,
		Status: status,
	}
	e.caster.Submit(n)
}

// Temp Global Event testing struct
var EV = NewEvents()

// Insert random events for testing
func (e *Events) InsertRandom() {
	for {
		wait := time.Duration(rand.Intn(4))
		logger.Critical("%d", wait)
		time.Sleep(wait * time.Second)
		e.Insert("alert", time.Now().String())
	}
}
