package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")

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
        "@googleapis//:google_api_annotations_proto",
        "@googleapis//:google_api_http_proto",
        "@googleapis//:google_longrunning_operations_proto",
        "@googleapis//:google_rpc_status_proto",
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
