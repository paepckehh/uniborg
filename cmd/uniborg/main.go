package main

import (
	"fmt"
	"time"
	// "paepcke.de/uniborg"
)

const (
	_app     = "[UNIBORG-CLI]"
	_version = "[v0.0.1]"
)

func main() {

	// Startup
	t0 := time.Now()
	fmt.Println(_app + "[STARTUP]" + _version)

	// Finish
	fmt.Println(_app + "[END][RUNTIME]:" + time.Now().Sub(t0).String())
}
