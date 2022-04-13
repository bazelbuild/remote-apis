package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@bazel_remote_apis_imports//:imports.bzl", "cc_grpc_library")

licenses(["notice"])

proto_library(
    name = "remote_asset_proto",
    srcs = ["remote_asset.proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@googleapis//:google_api_annotations_proto",
        "@googleapis//:google_api_http_proto",
        "@googleapis//:google_rpc_status_proto",
    ],
)

alias(
    name = "remote_asset_java_proto",
    actual = "//build/bazel/remote/asset/v1/java:remote_asset_java_proto",
)

cc_grpc_library(
    name = "remote_asset_cc_proto",
    srcs = ["remote_asset.proto"],
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_cc_proto",
        "@googleapis//:google_api_annotations_cc_proto",
        "@googleapis//:google_api_http_cc_proto",
        "@googleapis//:google_rpc_status_cc_proto",
    ],
)

alias(
    name = "remote_asset_go_proto",
    actual = "//build/bazel/remote/asset/v1/go:remote_asset_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/remote/asset/v1/go:go_default_library",
)
