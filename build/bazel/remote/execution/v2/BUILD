# gazelle:ignore

load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

licenses(["notice"])

proto_library(
    name = "remote_execution_proto",
    srcs = ["remote_execution.proto"],
    deps = [
        "//build/bazel/semver:semver_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
        "@com_google_googleapis//google/api:annotations_proto",
        "@com_google_googleapis//google/api:http_proto",
        "@com_google_googleapis//google/longrunning:operations_proto",
        "@com_google_googleapis//google/rpc:status_proto",
    ],
)

alias(
    name = "remote_execution_java_proto",
    actual = "//build/bazel/remote/execution/v2/java:remote_execution_java_proto",
)

alias(
    name = "remote_execution_cc_proto",
    actual = "//build/bazel/remote/execution/v2/cc:remote_execution_cc_grpc",
)

alias(
    name = "remote_execution_go_proto",
    actual = "//build/bazel/remote/execution/v2/go:remote_execution_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/remote/execution/v2/go:go_default_library",
)
