package goactor

import (
	"log"
)

type Any interface{}

type ActorInterface interface {
	Inbox() chan Any
	Act(message Any)
}

type Actor struct {
	inbox chan Any
}

func (this Actor) Inbox() chan Any {
	return this.inbox
}

func NewActor() Actor {
	return Actor{
		inbox: make(chan Any),
	}
}

func Go(actor ActorInterface, name string) {
	go func() {
		for {
			message, ok := <-actor.Inbox()
			if !ok {
				log.Printf("[%s] Inbox is unreachable", name)
				break
			}

			actor.Act(message)
		}
	}()
}

func (actor *Actor) Send(message Any) {
	go func() {
		actor.SyncSend(message)
	}()
}

func (actor *Actor) SyncSend(message Any) {
	actor.Inbox() <- message
}

func (actor *Actor) Die() {
	close(actor.Inbox())
}
