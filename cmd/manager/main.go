package main

import (
	"github.com/deissh/lambda/pkg/manager"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	m := manager.Create()

	all, err := m.GetActive()
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range all {
		_ = m.Stop(container.ID)
		log.Println(container.ID, " | ", container.State)
	}
}
