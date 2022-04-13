package(default_visibility = ["//build/bazel/remote/execution/v2:__pkg__"])

load("//internal:cc_grpc_library.bzl", "cc_grpc_codegen")

cc_proto_library(
    name = "remote_execution_cc_proto",
    deps = ["//build/bazel/remote/execution/v2:remote_execution_proto"],
)

cc_grpc_codegen(
    name = "remote_execution_cc_grpc_codegen",
    proto = "//build/bazel/remote/execution/v2:remote_execution_proto",
)

cc_library(
    name = "remote_execution_cc_grpc",
    srcs = [":remote_execution_cc_grpc_codegen"],
    hdrs = [":remote_execution_cc_grpc_codegen"],
    include_prefix = "build/bazel/remote/execution/v2",
    strip_include_prefix = "/" + package_name(),
    deps = [
        ":remote_execution_cc_proto",
        "@com_github_grpc_grpc//:grpc++_codegen_proto",
    ],
)
