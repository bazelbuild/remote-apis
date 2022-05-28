"""Repository rules and macros which are expected to be called from WORKSPACE
file of either bazel_remote_apis itself or any third_party repository which
consumes bazel_remote_apis as its dependency.

This is adapted from
https://github.com/googleapis/googleapis/blob/master/repository_rules.bzl
"""

def _switched_rules_impl(ctx):
    ctx.file("BUILD.bazel", "")
    ctx.file("imports.bzl", "\n".join(ctx.attr.lines))

switched_rules = repository_rule(
    implementation = _switched_rules_impl,
    attrs = {
        "lines": attr.string_list(
            allow_empty = True,
            mandatory = False,
        ),
    },
)

def switched_rules_by_language(
        name,
        java = False,
        go = False,
        cc = False,
        rules_override = {}):
    """Switches rules in the generated imports.bzl between no-op and the actual
    implementation.

    This defines which language-specific rules should be enabled during the
    build. All non-enabled language-specific rules will default to no-op
    implementations.

    For example, to use this rule and enable Java and Go rules, add the
    following in the external repository which imports bazel_remote_apis
    repository and its corresponding dependencies:

        load("@bazel_remote_apis//:repository_rules.bzl", "switched_rules_by_language")

        switched_rules_by_language(
            name = "bazel_remote_apis_imports",
            go = True,
            java = True,
        )

    Then import e.g. "go_library" from @bazel_remote_apis_imports in your BUILD files:

        load("@bazel_remote_apis_imports//:imports.bzl", "go_library")

    Note, for build to work you must also import the language-specific transitive dependencies.

    Args:
        name (str): name of a target, is expected to be "bazel_remote_apis_imports".
        java (bool): Enable Java specific rules. False by default.
        go (bool): Enable Go specific rules. False by default.
        cc (bool): Enable C++ specific rules. False by default.
        rules_override (dict): Custom rule overrides (for advanced usage).
    """

    loads = {}
    symbols = {}

    symbols["java_enabled"] = java
    if java:
        loads["java_proto_library"] = "@rules_java//java:defs.bzl"
        symbols["java_proto_library"] = "_java_proto_library"

    symbols["go_enabled"] = go
    if go:
        loads["go_proto_library"] = "@io_bazel_rules_go//proto:def.bzl"
        symbols["go_proto_library"] = "_go_proto_library"

    symbols["cc_enabled"] = cc
    if cc:
        loads["cc_grpc_library"] = "@com_github_grpc_grpc//bazel:cc_grpc_library.bzl"
        symbols["cc_grpc_library"] = "_cc_grpc_library"

    extra_lines = []
    for k, v in rules_override:
        if not v:
            if k in loads:
                loads.pop(k)
            if k in symbols:
                symbols.pop(k)
        elif v.startswith("@"):
            loads[k] = v
            symbols[k] = "_" + k
        elif v.startswith("native."):
            if k in loads:
                loads.pop(k)
            loads[k] = v
        else:
            extra_lines.append(v)

    load_lines = [_load(v, k) for k, v in loads.items()]
    symbol_lines = ["{} = {}".format(k, v) for k, v in symbols.items()]
    switched_rules(
        name = name,
        lines = load_lines + [
            "def maybe_alias(name, actual, enabled, visibility = None):",
            "    if enabled:",
            "        native.alias(",
            "            name = name,",
            "            actual = actual,",
            "            visibility = visibility,",
            "        )",
            "",
            "# Export symbols",
        ] + symbol_lines + [
            "",
            "# Extra statements from rules_override",
        ] + extra_lines,
    )

def _load(pkg, symbol):
    return "load('{0}', _{1} = '{1}')".format(pkg, symbol)
