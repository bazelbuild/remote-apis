package main

import (
	"fmt"

	rapb "github.com/bazelbuild/remote-apis/build/bazel/remote/asset/v1"
	repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	lspb "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1"
)

func main() {
	_ = repb.NewExecutionClient(nil)
	_ = rapb.NewPushClient(nil)
	_ = lspb.NewLogStreamServiceClient(nil)

	fmt.Println("Go example compiles and runs!")
}
