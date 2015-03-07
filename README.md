# goactor

Thin Actor implementation in golang.

## Installation

Specify it as a dependency:

```go
import "github.com/waterlink/goactor"
```

And run

```bash
go get
```

## Usage

Simple example

```go
type Relationships struct {
  Inbox chan Event,
}

func (this Relationships) Actor() {
  event, ok := <-this.Inbox
  if !ok {
    this.Die("Inbox is unreachable")
  }

  # ... handle event ...
}
```

To run this actor:

```go
relationships = Relationships{
  Inbox: make(chan Event),
}
goactor.Go(relationships)
```

To send anything to its inbox, one can use:

```go
relationships.Send(anEvent)
```

Actor needs following methods to be implemented:

- `(goactor.Actor) Actor()` - one lifecylce: get inbox message and do something important
- `(goactor.Actor) Inbox` of type `chan interface{}`

## Contributing

1. Fork it ( https://github.com/waterlink/goactor/fork )
2. Create a branch ( `git checkout -b a-feature` )
3. Commit your changes ( `git commit -am "A feature"` )
4. Push to the branch ( `git push -u origin a-feature` )
5. Create a new Pull Request
