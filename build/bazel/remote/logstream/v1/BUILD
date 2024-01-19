load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_go//proto:def.bzl", "go_grpc_library")
load("@rules_proto_grpc_cpp//:defs.bzl", "cpp_grpc_library")

# gazelle:ignore
package(default_visibility = ["//visibility:public"])

licenses(["notice"])

proto_library(
    name = "remote_logstream_proto",
    srcs = ["remote_logstream.proto"],
)

go_grpc_library(
    name = "remote_logstream_go_grpc",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1",
    protos = [":remote_logstream_proto"],
)

cpp_grpc_library(
    name = "remote_logstream_cpp_grpc",
    protos = [":remote_logstream_proto"],
)
