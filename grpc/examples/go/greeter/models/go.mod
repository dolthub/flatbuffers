module github.com/google/flatbuffers/grpc/examples/go/greeter/models

go 1.15

replace github.com/dolthub/flatbuffers/v23 => ../../../../..

require (
	github.com/dolthub/flatbuffers/v23 v23.3.3
	google.golang.org/grpc v1.35.0
)
