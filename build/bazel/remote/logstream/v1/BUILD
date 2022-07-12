package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")

licenses(["notice"])

proto_library(
    name = "remote_logstream_proto",
    srcs = ["remote_logstream.proto"],
)

alias(
    name = "remote_logstream_java_proto",
    actual = "//build/bazel/remote/logstream/v1/java:remote_logstream_java_proto",
)

alias(
    name = "remote_logstream_cc_proto",
    actual = "//build/bazel/remote/logstream/v1/cc:remote_logstream_cc_grpc",
)

alias(
    name = "remote_logstream_go_proto",
    actual = "//build/bazel/remote/logstream/v1/go:remote_logstream_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/remote/logstream/v1/go:go_default_library",
)
