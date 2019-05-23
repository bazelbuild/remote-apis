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
    strip_prefix = "grpc-1.21.0",
    urls = [
        "https://github.com/grpc/grpc/archive/v1.21.0.tar.gz",
        "https://mirror.bazel.build/github.com/grpc/grpc/archive/v1.21.0.tar.gz",
    ],
    sha256 = "8da7f32cc8978010d2060d740362748441b81a34e5425e108596d3fcd63a97f2",
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

# We have to import zlib directly ourselves, because protobuf_deps.bzl isn't
# part of the protobuf release yet
# (https://github.com/protocolbuffers/protobuf/issues/5918). This should be
# fixed in 3.8.0.
http_archive(
    name = "net_zlib",
    build_file = "@com_google_protobuf//:third_party/zlib.BUILD",
    sha256 = "6d4d6640ca3121620995ee255945161821218752b551a1a180f4215f7d124d45",
    strip_prefix = "zlib-1.2.11",
    urls = ["https://zlib.net/zlib-1.2.11.tar.gz"],
)

# Pull in go rules.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "86ae934bd4c43b99893fc64be9d9fc684b81461581df7ea8fc291c816f5ee8c5",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.3/rules_go-0.18.3.tar.gz"],
)
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
