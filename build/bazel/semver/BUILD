package(default_visibility = ["//visibility:public"])
load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

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

go_proto_library(
    name = "semver_go_proto",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
    proto = ":semver_proto",
)

go_library(
    name = "go_default_library",
    embed = [":semver_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
)
