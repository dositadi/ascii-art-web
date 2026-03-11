package main

import (
	a "ascii-web/internal/config"
)

// This is the engine the sets the system in motion
func main() {
	app := a.App{}

	app.Run()
}
