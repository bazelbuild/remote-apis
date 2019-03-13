package(default_visibility = ["//visibility:public"])
load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")

licenses(["notice"])

proto_library(
    name = "semver_proto",
    srcs = ["semver.proto"],
)

java_proto_library(
    name = "semver_java_proto",
    deps = [":semver_proto"],
)

cc_grpc_library(
    name = "semver_cc_proto",
    srcs = ["semver.proto"],
    deps = [],
    proto_only = False,
    well_known_protos = True,
    use_external = False,
)

