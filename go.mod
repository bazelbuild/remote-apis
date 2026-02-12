module github.com/bazelbuild/remote-apis

go 1.21
toolchain go1.24.1

require (
	cloud.google.com/go/longrunning v0.5.12
	google.golang.org/genproto/googleapis/api v0.0.0-20240812133136-8ffd90a71988
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240812133136-8ffd90a71988
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
