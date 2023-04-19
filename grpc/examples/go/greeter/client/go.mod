module github.com/google/flatbuffers/grpc/examples/go/greeter/client

go 1.15

replace github.com/google/flatbuffers/grpc/examples/go/greeter/models v0.0.0 => ../models

replace github.com/dolthub/flatbuffers/v23 => ../../../../..

require (
	github.com/google/flatbuffers/grpc/examples/go/greeter/models v0.0.0
	github.com/dolthub/flatbuffers/v23 v23.3.3
	google.golang.org/grpc v1.35.0
)
