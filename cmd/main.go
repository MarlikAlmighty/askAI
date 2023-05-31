package main

import (
	"github.com/MarlikAlmighty/kickHisAss/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("[Fatal]: start bot %v\n", err)
	}
}
