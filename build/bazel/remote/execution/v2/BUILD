package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@bazel_remote_apis_imports//:imports.bzl", "cc_grpc_library")

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

cc_grpc_library(
    name = "remote_execution_cc_proto",
    srcs = ["remote_execution.proto"],
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = [
        "//build/bazel/semver:semver_cc_proto",
        "@googleapis//:google_api_annotations_cc_proto",
        "@googleapis//:google_api_http_cc_proto",
        "@googleapis//:google_longrunning_operations_cc_proto",
        "@googleapis//:google_rpc_status_cc_proto",
    ],
)

alias(
    name = "remote_execution_go_proto",
    actual = "//build/bazel/remote/execution/v2/go:remote_execution_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/remote/execution/v2/go:go_default_library",
)
