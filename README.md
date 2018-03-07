# `gqlgen` Interface Bug Reproduction

To reproduce:

- Put this in your `$GOPATH`:
  - `git clone https://github.com/ereyes01/gqlgen-interface-bug.git $GOPATH/src/github.com/ereyes01/gqlgen-interface-bug`
- Install dependencies:
  - `go get github.com/gorilla/websocket`
  - `go get github.com/vektah/gqlgen`
- Generate the graphql code:
  - `cd $GOPATH/src/github.com/ereyes01/gqlgen-interface-bug/shapes && gqlgen -out generated.go -package shapes -typemap types.json schema.graphql`
- Try to build the playground server:
  - `go install github.com/ereyes01/gqlgen-interface-bug/shapes/server`

The last command will fail with the following:

```
# github.com/ereyes01/gqlgen-interface-bug/shapes
shapes/generated.go:711:2: impossible type switch case: *obj (type Shape) cannot have dynamic type Circle (Area method has pointer receiver)
shapes/generated.go:716:2: impossible type switch case: *obj (type Shape) cannot have dynamic type Rectangle (Area method has pointer receiver)
```

