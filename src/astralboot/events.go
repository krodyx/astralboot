package main

// A channel muxer for the event system

import (
	//	"fmt"
	"github.com/dustin/go-broadcast"
)

type Events struct {
	caster broadcast.Broadcaster
}

func NewEvents() (e *Events) {
	e = &Events{}
	e.caster = broadcast.NewBroadcaster(1024)
	return
}

type notif struct {
	Name   string
	Status string
}
