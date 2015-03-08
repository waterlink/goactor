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
  goactor.Actor
}

func (this *Relationships) Act(message goactor.Any) {
  event, ok := message.(Event)
  if !ok {
    return     # ignore or handle it
  }

  # ... handle event ...
}
```

To run this actor:

```go
relationships := Relationships{goactor.NewActor()}
goactor.Go(relationships, "Relationships Task")
```

To send anything to its inbox, one can use:

```go
goactor.Send(relationships, anEvent)
```

Actor needs following methods to be implemented:

- `(goactor.Actor) Act(message Any)` - one lifecylce: get inbox message and do something important

Look at the full [example](examples/example.go)

For further details you can look at the test: [goactor_test.go](goactor_test.go)

## Contributing

1. Fork it ( https://github.com/waterlink/goactor/fork )
2. Create a branch ( `git checkout -b a-feature` )
3. Commit your changes ( `git commit -am "A feature"` )
4. Push to the branch ( `git push -u origin a-feature` )
5. Create a new Pull Request
