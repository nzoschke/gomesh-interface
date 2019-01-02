# GoMesh Interface

This repo contains generated client and server interfaces for [GoMesh](https://github.com/nzoschke/gomesh). We keep these in their own repo to isolate API design and definitions from generated code. This facilitates a development flow of:

1. Design APIs as .proto definitions ([gomesh-proto](https://github.com/nzoschke/gomesh-proto))
2. Generate Go server interfaces ([gomesh-interface](https://github.com/nzoschke/gomesh-interface))
3. Build gRPC services that satisfy the server interfaces ([gomesh](https://github.com/nzoschke/gomesh))

## Motivation

Building a reliable service oriented architecture (SOA) is easier than ever, once you learn the gRPC framework and ecosystem of tools that interoperate around Protocol Buffer service definitions.

[GoMesh](https://github.com/nzoschke/gomesh) is an example SOA with all of the proto and gRPC services configured correctly and explaiend in depth. See the [gomesh/docs folder](https://github.com/nzoschke/gomesh/tree/master/docs) for detailed guides about service definitions, proxies, remote procedure calls, API gateways, data stores, observability, versioning and more.

With this foundation you can skip over all the setup, and focus entirely on your business logic code.

## Quick Start

This project uses:

- [Go 1.11](https://golang.org/)
- [grpc-go](https://github.com/grpc/grpc-go)

### Install the tools

```console
$ go get -u github.com/nzoschke/gomesh-interface/gen/go/users/v1/...
$ go get -u google.golang.org/grpc
```

### Use a client

We start by importing the `users/v1` Go package in our program, and using its strongly typed client methods, and request and response structs:

```go
package main

import (
	"context"
	"log"

	users "github.com/nzoschke/gomesh-interface/gen/go/users/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := users.NewUsersClient(conn)
	ctx := context.Background()

	r, err := c.Create(ctx, &users.CreateRequest{
		Parent: "mycorp",
		User: &users.User{
			DisplayName: "My Name",
			Name:        "myusername",
		},
	})

	log.Printf("User created at %s", r.CreateTime)
}
```

This gives us confidence in our generated Go package.

## Docs

Check out [gomesh/docs folder](https://github.com/nzoschke/gomesh/tree/master/docs) where each component is explained in more detail.

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/nzoschke/gomesh-interface/issues).

## License

Apache 2.0 © 2019 Noah Zoschke