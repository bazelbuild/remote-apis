package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")

licenses(["notice"])

proto_library(
    name = "semver_proto",
    srcs = ["semver.proto"],
)

alias(
    name = "semver_java_proto",
    actual = "//build/bazel/semver/java:semver_java_proto",
)

alias(
    name = "semver_cc_proto",
    actual = "//build/bazel/semver/cc:semver_cc_proto",
)

alias(
    name = "semver_go_proto",
    actual = "//build/bazel/semver/go:semver_go_proto",
)

alias(
    name = "go_default_library",
    actual = "//build/bazel/semver/go:go_default_library",
)
