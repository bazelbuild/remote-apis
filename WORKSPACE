workspace(name = "bazel_remote_apis")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains(version = "1.15.6")

#http_archive(
#    name = "upb",
#    sha256 = "72d86e89d0345cba15106dfb5fc5f550dc3de0f1871599b8a481bdda54e7dee0"
#    strip_prefix = "upb-0f40d59258173b13f989a9f801967f44291fa30d",
#    urls = ["https://github.com/protocolbuffers/upb/archive/0f40d59258173b13f989a9f801967f44291fa30d.zip"],
#)
#
#load("@upb//bazel:workspace_deps.bzl", "upb_deps")
#
#upb_deps()

http_archive(
    name = "bazel_skylib",
    sha256 = "97e70364e9249702246c0e9444bccdc4b847bed1eb03c5a3ece4f83dfe6abc44",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

# Gazelle, which we need in order to selectively pull in Gazelle dependencies for Go.
http_archive(
    name = "bazel_gazelle",
    sha256 = "41bff2a0b32b02f20c227d234aa25ef3783998e5453f7eade929704dcff7cd4b",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.0/bazel-gazelle-v0.19.0.tar.gz"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# Needed for protobuf.
http_archive(
    name = "com_google_protobuf",
    sha256 = "678d91d8a939a1ef9cb268e1f20c14cd55e40361dc397bb5881e4e1e532679b1",
    strip_prefix = "protobuf-3.10.1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.10.1.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# Needed for C++ gRPC.
http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "b391a327429279f6f29b9ae7e5317cd80d5e9d49cc100e6d682221af73d984a6",
    strip_prefix = "grpc-93e8830070e9afcbaa992c75817009ee3f4b63a0",  # v1.24.3 with fixes
    urls = ["https://github.com/grpc/grpc/archive/93e8830070e9afcbaa992c75817009ee3f4b63a0.zip"],
)

# Pull in all gRPC dependencies.
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

# More gRPC dependencies. grpc_extra_deps does not work out of the box.
#load("@upb//bazel:workspace_deps.bzl", "upb_deps")
load("@build_bazel_rules_apple//apple:repositories.bzl", "apple_rules_dependencies")
load("@build_bazel_apple_support//lib:repositories.bzl", "apple_support_dependencies")

#upb_deps()

apple_rules_dependencies()

apple_support_dependencies()

load("@upb//bazel:repository_defs.bzl", "bazel_version_repository")

bazel_version_repository(
    name = "bazel_version",
)

bind(
    name = "grpc_cpp_plugin",
    actual = "@com_github_grpc_grpc//:grpc_cpp_plugin",
)

bind(
    name = "grpc_lib",
    actual = "@com_github_grpc_grpc//:grpc++",
)

# Needed for the googleapis protos.
http_archive(
    name = "googleapis",
    sha256 = "825be3bf70632d273a10e41db1f458211446d129063c4f29cee854f9bcbf834b",
    strip_prefix = "googleapis-61ab0348bd228c942898aee291d677f0afdb888c",
    urls = ["https://github.com/googleapis/googleapis/archive/61ab0348bd228c942898aee291d677f0afdb888c.zip"],
)

load("@googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    go = True,
    java = True,
)
