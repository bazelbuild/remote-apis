load("@rules_go//proto:def.bzl", "go_grpc_library")
load("@rules_java//java:defs.bzl", "java_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc_cpp//:defs.bzl", "cpp_grpc_library")

# gazelle:ignore
package(default_visibility = ["//visibility:public"])

licenses(["notice"])

proto_library(
    name = "remote_asset_proto",
    srcs = ["remote_asset.proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_googleapis//google/api:annotations_proto",
        "@com_google_googleapis//google/api:http_proto",
        "@com_google_googleapis//google/rpc:status_proto",
    ],
)

go_grpc_library(
    name = "remote_asset_go_grpc",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/asset/v1",
    protos = [":remote_asset_proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_go_grpc",
        "@org_golang_google_genproto_googleapis_api//annotations",
        "@org_golang_google_genproto_googleapis_rpc//status",
    ],
)  # keep

alias(
    name = "go_default_library",
    actual = ":remote_asset_go_grpc",
)

cpp_grpc_library(
    name = "remote_asset_cpp_grpc",
    protos = [":remote_asset_proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_cpp_grpc",
    ],
)

# TODO: Include/replace with gRPC stub generation targets when feasible
java_proto_library(
    name = "remote_asset_java_proto",
    deps = [":remote_asset_proto"],
)
