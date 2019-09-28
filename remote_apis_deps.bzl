"""Load dependencies needed to depend on the RE API repo."""

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

def remote_apis_go_deps():
    """Load dependencies needed to depend on RE API for Go"""
    go_rules_dependencies()
    go_register_toolchains()
    gazelle_dependencies()
    if "com_github_golang_protobuf" not in native.existing_rules():
        go_repository(
            name = "com_github_golang_protobuf",
            importpath = "github.com/golang/protobuf",
            tag = "v1.3.2",
        )
    if "org_golang_google_grpc" not in native.existing_rules():
        go_repository(
            name = "org_golang_google_grpc",
            build_file_proto_mode = "disable",
            importpath = "google.golang.org/grpc",
            sum = "h1:J0UbZOIrCAl+fpTOf8YLs4dJo8L/owV4LYVtAXQoPkw=",
            version = "v1.22.0",
        )
    if "org_golang_x_net" not in native.existing_rules():
        go_repository(
            name = "org_golang_x_net",
            importpath = "golang.org/x/net",
            sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
            version = "v0.0.0-20190311183353-d8887717615a",
        )
    if "org_golang_x_text" not in native.existing_rules():
        go_repository(
            name = "org_golang_x_text",
            importpath = "golang.org/x/text",
            sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
            version = "v0.3.0",
        )
