package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@bazel_remote_apis_imports//:imports.bzl", "cc_grpc_library")

licenses(["notice"])

proto_library(
    name = "semver_proto",
    srcs = ["semver.proto"],
)

alias(
    name = "semver_java_proto",
    actual = "//build/bazel/semver/java:semver_java_proto",
)

cc_grpc_library(
    name = "semver_cc_proto",
    srcs = ["semver.proto"],
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = [],
)

alias(
    name = "semver_go_proto",
    actual = "//build/bazel/semver/go:semver_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/semver/go:go_default_library",
)
