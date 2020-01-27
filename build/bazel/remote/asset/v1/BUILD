package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@bazel_remote_apis_imports//:imports.bzl", "java_proto_library", "go_library", "go_proto_library", "cc_grpc_library")

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

java_proto_library(
    name = "remote_asset_java_proto",
    deps = [":remote_asset_proto"],
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

go_proto_library(
    name = "remote_asset_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/asset/v1",
    proto = ":remote_asset_proto",
    deps = [
        "//build/bazel/remote/execution/v2:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":remote_asset_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/asset/v1",
)
