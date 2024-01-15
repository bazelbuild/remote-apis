load("@rules_go//go:def.bzl", "go_library")
load("@rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//build/bazel/semver:__pkg__"])

go_proto_library(
    name = "semver_go_proto",
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
    proto = "//build/bazel/semver:semver_proto",
)

go_library(
    name = "go_default_library",
    embed = [":semver_go_proto"],
    importpath = "github.com/bazelbuild/remote-apis/build/bazel/semver",
)
