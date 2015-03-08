package main

import (
	"github.com/waterlink/goactor"
	"log"
	"time"
)

type Event struct {
	Sender  string
	Message string
}

type Relationships struct {
	goactor.Actor
}

func (this *Relationships) Act(message goactor.Any) {
	event, _ := message.(Event)
	log.Print(event)
}

func main() {
	relationships := Relationships{goactor.NewActor()}
	goactor.Go(relationships, "Relationships Task")

	anEvent := Event{"the sender", "the message"}
	goactor.Send(relationships, anEvent)

	time.Sleep(50 * time.Millisecond)
}
