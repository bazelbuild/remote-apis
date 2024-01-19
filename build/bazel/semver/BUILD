# gazelle:ignore
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_go//proto:def.bzl", "go_grpc_library")

package(default_visibility = ["//visibility:public"])

licenses(["notice"])

proto_library(
    name = "semver_proto",
    srcs = ["semver.proto"],
)

go_grpc_library(
    name = "semver_go_grpc",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
    protos = [":semver_proto"],
)
