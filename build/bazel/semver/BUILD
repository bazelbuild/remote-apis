load("@bazel_remote_apis_imports//:imports.bzl", "cc_grpc_library", "go_enabled", "go_proto_library", "java_proto_library", "maybe_alias")
load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

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
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = [],
)

go_proto_library(
    name = "semver_go_proto",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
    proto = ":semver_proto",
)

maybe_alias(
    name = "go_default_library",
    actual = ":semver_go_proto",
    enabled = go_enabled,
)
