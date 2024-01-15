load("@rules_proto//proto:defs.bzl", "proto_library")

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
