# gazelle:ignore

load("@rules_go//proto:def.bzl", "go_grpc_library")
load("@rules_java//java:defs.bzl", "java_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc_cpp//:defs.bzl", "cpp_grpc_library")

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

go_grpc_library(
    name = "remote_execution_go_grpc",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2",
    protos = [":remote_execution_proto"],
    deps = [
        "//build/bazel/semver:semver_go_grpc",
        "@org_golang_google_genproto//googleapis/longrunning",
        "@org_golang_google_genproto_googleapis_api//annotations",
        "@org_golang_google_genproto_googleapis_rpc//http",
        "@org_golang_google_genproto_googleapis_rpc//status",
    ],
)

alias(
    name = "go_default_library",
    actual = ":remote_execution_go_grpc",
)

cpp_grpc_library(
    name = "remote_execution_cpp_grpc",
    protos = [":remote_execution_proto"],
    deps = [
        "//build/bazel/semver:annotations_cpp_grpc",
        "//build/bazel/semver:http_cpp_grpc",
        "//build/bazel/semver:operations_cpp_grpc",
        "//build/bazel/semver:semver_cpp_grpc",
    ],
)

java_proto_library(
    name = "remote_execution_java_grpc",
    deps = [":remote_execution_proto"],
)
