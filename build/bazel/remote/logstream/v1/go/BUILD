package(default_visibility = ["//build/bazel/remote/logstream/v1:__pkg__"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "remote_logstream_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1",
    proto = "//build/bazel/remote/logstream/v1:remote_logstream_proto",
)

go_library(
    name = "go_default_library",
    embed = [":remote_logstream_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1",
)
