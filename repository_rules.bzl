"""Deprecated stub for compatibility with previous releases."""

load("//:remote_apis_deps.bzl", "remote_apis_go_deps")

def switched_rules_by_language(
        name,
        java = False,
        go = False,
        cc = False,
        rules_override = {}):
    """Deprecated stub for compatibility with previous releases."""
    print(
        "The switched_rules_by_language macro is deprecated. Consumers of" +
        " @bazel_remote_apis should specify per-language dependencies in" +
        " their own workspace.",
    )
    if go:
        remote_apis_go_deps()
