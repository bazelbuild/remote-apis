workspace(name = "bazel_remote_apis")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Needed for protobuf.
http_archive(
    name = "bazel_skylib",
    # Commit f83cb8dd6f5658bc574ccd873e25197055265d1c of 2018-11-26
    sha256 = "ba5d15ca230efca96320085d8e4d58da826d1f81b444ef8afccd8b23e0799b52",
    strip_prefix = "bazel-skylib-f83cb8dd6f5658bc574ccd873e25197055265d1c",
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/archive/f83cb8dd6f5658bc574ccd873e25197055265d1c.tar.gz",
    ],
)

# Needed for "well-known protos" and protoc.
http_archive(
    name = "com_google_protobuf",
    sha256 = "3e933375ecc58d01e52705479b82f155aea2d02cc55d833f8773213e74f88363",
    strip_prefix = "protobuf-3.7.0",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protobuf-all-3.7.0.tar.gz"],
)

# Needed for the googleapis protos.
http_archive(
    name = "googleapis",
    build_file = "BUILD.googleapis",
    sha256 = "7b6ea252f0b8fb5cd722f45feb83e115b689909bbb6a393a873b6cbad4ceae1d",
    strip_prefix = "googleapis-143084a2624b6591ee1f9d23e7f5241856642f4d",
    urls = ["https://github.com/googleapis/googleapis/archive/143084a2624b6591ee1f9d23e7f5241856642f4d.zip"],
)

# Needed for C++ gRPC.
http_archive(
    name = "com_github_grpc_grpc",
    strip_prefix = "grpc-1.17.2",
    urls = [
        "https://github.com/grpc/grpc/archive/v1.17.2.tar.gz",
        "https://mirror.bazel.build/github.com/grpc/grpc/archive/v1.17.2.tar.gz",
    ],
    sha256 = "34ed95b727e7c6fcbf85e5eb422e962788e21707b712fdb4caf931553c2c6dbc",
)

# Pull in all gRPC dependencies.
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")
grpc_deps()

bind(
    name = "grpc_cpp_plugin",
    actual = "@com_github_grpc_grpc//:grpc_cpp_plugin",
)

bind(
    name = "grpc_lib",
    actual = "@com_github_grpc_grpc//:grpc++",
)
