package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@bazel_remote_apis_imports//:imports.bzl", "java_proto_library", "go_library", "go_proto_library", "cc_grpc_library")

licenses(["notice"])

proto_library(
    name = "remote_output_proto",
    srcs = ["remote_output.proto"],
    deps = [
        "//build/bazel/remote/execution/v2:remote_execution_proto",
        "@com_google_protobuf//:empty_proto",
    ],
)

java_proto_library(
    name = "remote_output_java_proto",
    deps = [":remote_output_proto"],
)

cc_grpc_library(
    name = "remote_output_cc_proto",
    srcs = ["remote_output.proto"],
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = ["//build/bazel/remote/execution/v2:remote_execution_cc_proto"],
)

go_proto_library(
    name = "remote_output_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/output/v1",
    proto = ":remote_output_proto",
    deps = ["//build/bazel/remote/execution/v2:go_default_library"],
)

go_library(
    name = "go_default_library",
    embed = [":remote_output_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/output/v1",
)
