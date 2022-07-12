package(default_visibility = ["//build/bazel/remote/logstream/v1:__pkg__"])

load("//internal:cc_grpc_library.bzl", "cc_grpc_codegen")

cc_proto_library(
    name = "remote_logstream_cc_proto",
    deps = ["//build/bazel/remote/logstream/v1:remote_logstream_proto"],
)

cc_grpc_codegen(
    name = "remote_logstream_cc_grpc_codegen",
    proto = "//build/bazel/remote/logstream/v1:remote_logstream_proto",
)

cc_library(
    name = "remote_logstream_cc_grpc",
    srcs = [":remote_logstream_cc_grpc_codegen"],
    hdrs = [":remote_logstream_cc_grpc_codegen"],
    include_prefix = "build/bazel/remote/logstream/v1",
    strip_include_prefix = "/" + package_name(),
    deps = [
        ":remote_logstream_cc_proto",
        "@com_github_grpc_grpc//:grpc++_codegen_proto",
    ],
)
