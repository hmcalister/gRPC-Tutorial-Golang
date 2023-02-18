# gRPC Tutorial - Golang
#### Author: Hayden McAlister

---

This repo is a first step into gRPC, using Golang for both server and client. Of course, in future this could be easily(?) extended to using other languages for the client and writing application code in Go.

The contents of this project are modified versions of [this excellent tutorial](https://www.golinuxcloud.com/golang-grpc/).

## Running the Project

In the first terminal, run:

```console
$ go run server/main.go
```

to start the server. A log message should be printed noting `gRPC server listening on [::]:50051`.

After this, in a separate terminal (with the server still running), run:

```console
$ go run client/main.go
```

There should be an immediate flurry of messages as new activities are processed by the server and the results printed by the client.

If an error occurs, the message should indicate what failed (registering the server port, dialing from the client, etc...). Good luck!

---
## Proto

`activity.proto` in the `proto` directory defines the gRPC protocol buffers for the client and server. We first define the syntax (`proto3` is the most recent, if omitted defaults to `proto2`), the package (under the `proto` directory), and the Go package we are working with. This information is important but more boilerplate than application specific.

We then define two messages: `NewActivity` and `Activity`. Messages are the payloads that can be passed between client and server when using services (see below). In this example, a client will eventually be able to pass the `NewActivity` message to the server, which will respond with an `Activity` message. This is a very basic example, but it shows that the client can define some data, pass it to a server, and get a response back. This could extend, for example, into a banking client passing a `TransferRequest` message to the central server, getting a `TransferReceipt` in return!

Finally, we define the `ActivityService` service. In gRPC, a service is a collection of possible requests the client can make to a server - in this example we say that the gRPC server will use the `ActivityService` and thus expose all the behaviors defined by it. We happen to only define one service method: `CreateActivity` which is a unary method. This means the client will send one request and get one response. [Other methods are available](https://grpc.io/docs/what-is-grpc/core-concepts/).

Once we have defined the protocol buffer code we must compile it for use in the client and server code. Because both the client and server are implemented in Golang, we can just compile the proto once, using

```console
$ protoc --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/activity.proto
```

Note we specify `go_out` to get Golang implementations of the proto code.

## Server

Our server is relatively basic, just implementing the `CreateActivity` method and setting up the listening connection and gRPC server (which is painless in Go). Note the `CreateActivity` implementation just takes the `NewActivity` struct and adds a single unique string to make the `Activity` struct for the proto response. Very boring!

## Client

Much like the server, the client code is very basic. We dial up the gRPC server, make a few requests for a new activity, and print the result. Yawn! Of course the real magic (and focus of this repo) is the gRPC communication between the client and server.