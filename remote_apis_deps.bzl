"""Load dependencies needed to depend on the RE API repo."""

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_download_sdk", "go_register_toolchains", "go_rules_dependencies")

def _maybe(repo_rule, name, **kwargs):
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)

def remote_apis_go_deps():
    """Load dependencies needed to depend on RE API for Go"""
    go_download_sdk(name = "go_sdk", version = "1.20.6")
    go_register_toolchains()

    # The version of this repo needs to be in sync with @googleapis
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:DTJM0R8LECCgFeUwApvcEJHz85HLagW8uRENYxHh1ww=",
        version = "v0.0.0-20240429193739-8cf5692501f6",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:DujSIu+2tC9Ht0aPNA7jgj23Iq8Ewi5sgkQ++wdvonE=",
        version = "v0.0.0-20240429193739-8cf5692501f6",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:WLbHekDbjK1fVFD3ibpFFVoyizlLRl73I7YKuAKilhU=",
        version = "v0.5.7",
    )

    go_rules_dependencies()
    gazelle_dependencies(go_sdk = "go_sdk")
    _maybe(
        go_repository,
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        tag = "v1.34.0",
    )
    _maybe(
        go_repository,
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        version = "v1.63.2",
    )
