package(default_visibility = ["//visibility:public"])
load("@bazel_remote_apis_imports//:imports.bzl", "java_proto_library", "go_library", "go_proto_library", "cc_grpc_library")

licenses(["notice"])

proto_library(
    name = "remote_execution_proto",
    srcs = ["remote_execution.proto"],
    deps = [
        "//build/bazel/semver:semver_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@googleapis//:google_api_annotations_proto",
        "@googleapis//:google_api_http_proto",
        "@googleapis//:google_longrunning_operations_proto",
        "@googleapis//:google_rpc_status_proto",
    ],
)

java_proto_library(
    name = "remote_execution_java_proto",
    deps = [":remote_execution_proto"],
)

cc_grpc_library(
    name = "remote_execution_cc_proto",
    srcs = ["remote_execution.proto"],
    deps = [
        "//build/bazel/semver:semver_cc_proto",
        "@googleapis//:google_api_annotations_cc_proto",
        "@googleapis//:google_api_http_cc_proto",
        "@googleapis//:google_longrunning_operations_cc_proto",
        "@googleapis//:google_rpc_status_cc_proto",
    ],
    proto_only = False,
    well_known_protos = True,
    use_external = False,
)

go_proto_library(
    name = "remote_execution_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2",
    proto = ":remote_execution_proto",
    deps = [
        "//build/bazel/semver:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/longrunning:longrunning_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":remote_execution_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2",
)
