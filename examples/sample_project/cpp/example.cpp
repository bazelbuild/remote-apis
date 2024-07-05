#include <iostream>

// #includes for these protos appear to be determined by the protobuf package,
// and not related to BUILD rule path in this repo.
#include "build/bazel/remote/asset/v1/remote_asset.grpc.pb.h"
#include "build/bazel/remote/asset/v1/remote_asset.pb.h"
#include "build/bazel/remote/execution/v2/remote_execution.grpc.pb.h"
#include "build/bazel/remote/execution/v2/remote_execution.pb.h"
#include "build/bazel/remote/logstream/v1/remote_logstream.grpc.pb.h"
#include "build/bazel/remote/logstream/v1/remote_logstream.pb.h"

int main(int argc, char **argv) {
  std::cout << "C++ example compiles and runs!" << std::endl;
}
