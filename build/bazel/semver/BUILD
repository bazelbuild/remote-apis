load("@rules_go//proto:def.bzl", "go_grpc_library")
load("@rules_java//java:defs.bzl", "java_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc_cpp//:defs.bzl", "cpp_grpc_library")

package(default_visibility = ["//visibility:public"])

licenses(["notice"])

proto_library(
    name = "semver_proto",
    srcs = ["semver.proto"],
)

# golang targets

go_grpc_library(
    name = "semver_go_grpc",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
    protos = [":semver_proto"],
)

alias(
    name = "go_default_library",
    actual = ":semver_go_grpc",
)

# C++ targets
cpp_grpc_library(
    name = "semver_cpp_grpc",
    protos = [":semver_proto"],
)

# java targets
# TODO: Include/replace with gRPC stub generation targets when feasible
java_proto_library(
    name = "semver_java_proto",
    deps = [":semver_proto"],
)

# TODO: Should the dependencies below be moved/removed in favor of a third-party
# target reference?

## protobuf proto libraries

cpp_grpc_library(
    name = "status_cpp_grpc",
    protos = ["@com_google_googleapis//google/rpc:status_proto"],
)

## googleapis proto libraries

cpp_grpc_library(
    name = "annotations_cpp_grpc",
    protos = ["@com_google_googleapis//google/api:annotations_proto"],
    deps = [
        ":http_cpp_grpc",
    ],
)

cpp_grpc_library(
    name = "client_cpp_grpc",
    protos = [
        "@com_google_googleapis//google/api:client_proto",
        "@com_google_googleapis//google/api:launch_stage_proto",
    ],
)

cpp_grpc_library(
    name = "http_cpp_grpc",
    protos = ["@com_google_googleapis//google/api:http_proto"],
)

cpp_grpc_library(
    name = "operations_cpp_grpc",
    protos = ["@com_google_googleapis//google/longrunning:operations_proto"],
    deps = [
        ":annotations_cpp_grpc",
        ":client_cpp_grpc",
        ":status_cpp_grpc",
    ],
)
