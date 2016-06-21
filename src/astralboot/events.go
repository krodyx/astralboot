package main

// A channel muxer for the event system

import (
	//	"fmt"
	"github.com/dustin/go-broadcast"
)

type Events struct {
	caster broadcast.Broadcaster
}
