package goactor

import (
	"fmt"
	"testing"
)

type AnActor struct {
	Actor
	outbox chan string
}

func (this *AnActor) Act(message Any) {
	response := fmt.Sprintf("Got '%s'", message)
	this.outbox <- response
}

func expectResponseToEq(t *testing.T, outbox chan string, expected string) {
	if actual := <-outbox; actual != expected {
		t.Errorf("Expected to receive '%s' but received '%s'", expected, actual)
	}
}

func TestHandlesInboxMessages(t *testing.T) {
	outbox := make(chan string)
	anActor := AnActor{
		Actor:  NewActor(),
		outbox: outbox,
	}
	Go(&anActor, "String Actor")

	anActor.Send("hello, world")
	expectResponseToEq(t, outbox, "Got 'hello, world'")

	anActor.Send("hello, goworld")
	expectResponseToEq(t, outbox, "Got 'hello, goworld'")
}

func TestClosedInbox(t *testing.T) {
	anActor := AnActor{
		Actor: NewActor(),
	}
	Go(&anActor, "String Actor")

	close(anActor.Inbox())
}

type AnIntegerActor struct {
	Actor
	outbox chan int
}

func (this *AnIntegerActor) Act(message Any) {
	if integerMessage, ok := message.(int); ok {
		response := integerMessage + 1
		this.outbox <- response
	}
}

func expectIntegerResponseToEq(t *testing.T, outbox chan int, expected int) {
	if actual := <-outbox; actual != expected {
		t.Errorf("Expected to receive '%d' but received '%d'", expected, actual)
	}
}

func TestWorksWithDifferentType(t *testing.T) {
	outbox := make(chan int)
	anActor := AnIntegerActor{
		Actor:  NewActor(),
		outbox: outbox,
	}
	Go(&anActor, "Integer Actor")

	anActor.Send(41)
	expectIntegerResponseToEq(t, outbox, 42)
}

func TestDie(t *testing.T) {
	anActor := &AnActor{
		Actor:  NewActor(),
		outbox: make(chan string),
	}
	Go(anActor, "Dying Actor")

	anActor.Die()

	_, ok := <-anActor.Inbox()
	if ok {
		t.Error("Expected Inbox to be closed")
	}
}
