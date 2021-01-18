package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "cc_grpc_library",
    "go_library",
    "go_proto_library",
    "java_proto_library",
)

licenses(["notice"])

proto_library(
    name = "remote_logstream_proto",
    srcs = ["remote_logstream.proto"],
    deps = [
    ],
)

java_proto_library(
    name = "remote_logstream_java_proto",
    deps = [":remote_logstream_proto"],
)

cc_grpc_library(
    name = "remote_logstream_cc_proto",
    srcs = ["remote_logstream.proto"],
    proto_only = False,
    use_external = False,
    well_known_protos = True,
    deps = [
    ],
)

#go_proto_library(
#    name = "remote_logstream_go_proto",
#    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
#    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1",
#    proto = ":remote_logstream_proto",
#    deps = [
#    ],
#)
#
#go_library(
#    name = "go_default_library",
#    embed = [":remote_logstream_go_proto"],
#    importpath = "github.com/bazelbuild/remote-apis/build/bazel/remote/logstream/v1",
#)
