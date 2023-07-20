package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")

licenses(["notice"])

proto_library(
    name = "remote_asset_proto",
    srcs = ["remote_asset.proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@googleapis//google/api:annotations_proto",
        "@googleapis//google/api:http_proto",
        "@googleapis//google/rpc:status_proto",
    ],
)

alias(
    name = "remote_asset_java_proto",
    actual = "//build/bazel/remote/asset/v1/java:remote_asset_java_proto",
)

alias(
    name = "remote_asset_cc_proto",
    actual = "//build/bazel/remote/asset/v1/cc:remote_asset_cc_grpc",
)

alias(
    name = "remote_asset_go_proto",
    actual = "//build/bazel/remote/asset/v1/go:remote_asset_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/remote/asset/v1/go:go_default_library",
)
