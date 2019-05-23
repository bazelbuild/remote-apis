package(default_visibility = ["//visibility:public"])
load("@bazel_remote_apis_imports//:imports.bzl", "java_proto_library", "go_library", "go_proto_library", "cc_grpc_library")

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
