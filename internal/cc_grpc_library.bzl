"""Rule for running the gRPC C++ code generator.

This is a simplified and modernised version of the upstream gRPC
`bazel/cc_grpc_library.bzl` file, which as of release v1.45
(published 2022-03-19) does not support separating the `proto_library` and
`cc_proto_library` targets into separate packages or repositories.

The following logic should eventually find a home in upstream gRPC, rules_proto,
or rules_cc so that the Bazel Remote APIs repository can be further decoupled
from language-specific concerns.
"""

load("@com_github_grpc_grpc//bazel:protobuf.bzl", "get_include_protoc_args")

_EXT_PROTO = ".proto"
_EXT_PROTODEVEL = ".protodevel"
_EXT_GRPC_HDR = ".grpc.pb.h"
_EXT_GRPC_SRC = ".grpc.pb.cc"

def _drop_proto_ext(name):
    if name.endswith(_EXT_PROTO):
        return name[:-len(_EXT_PROTO)]
    if name.endswith(_EXT_PROTODEVEL):
        return name[:-len(_EXT_PROTODEVEL)]
    fail("{!r} does not end with {!r} or {!r}".format(
        name,
        _EXT_PROTO,
        _EXT_PROTODEVEL,
    ))

def _proto_srcname(file):
    """Return the Protobuf source name for a proto_library source file.

    The source name is what the Protobuf compiler uses to identify a .proto
    source file. It is relative to the compiler's `--proto_path` flag.
    """
    ws_root = file.owner.workspace_root
    if ws_root != "" and file.path.startswith(ws_root):
        return file.path[len(ws_root) + 1:]
    return file.short_path

def _cc_grpc_codegen(ctx):
    """Run the gRPC C++ code generator to produce sources and headers"""
    proto = ctx.attr.proto[ProtoInfo]
    proto_srcs = proto.check_deps_sources.to_list()
    proto_imports = proto.transitive_imports.to_list()

    protoc_out = ctx.actions.declare_directory(ctx.attr.name + "_protoc_out")
    protoc_outputs = [protoc_out]
    rule_outputs = []

    for proto_src in proto_srcs:
        srcname = _drop_proto_ext(_proto_srcname(proto_src))
        basename = _drop_proto_ext(proto_src.basename)

        out_hdr = ctx.actions.declare_file(basename + _EXT_GRPC_HDR)
        out_src = ctx.actions.declare_file(basename + _EXT_GRPC_SRC)

        protoc_out_prefix = protoc_out.basename
        protoc_out_hdr = ctx.actions.declare_file(
            "{}/{}".format(protoc_out_prefix, srcname + _EXT_GRPC_HDR),
        )
        protoc_out_src = ctx.actions.declare_file(
            "{}/{}".format(protoc_out_prefix, srcname + _EXT_GRPC_SRC),
        )

        rule_outputs.extend([out_hdr, out_src])
        protoc_outputs.extend([protoc_out_hdr, protoc_out_src])

        ctx.actions.expand_template(
            template = protoc_out_hdr,
            output = out_hdr,
            substitutions = {},
        )
        ctx.actions.expand_template(
            template = protoc_out_src,
            output = out_src,
            substitutions = {},
        )

    plugin = ctx.executable._protoc_gen_grpc
    protoc_args = ctx.actions.args()
    protoc_args.add("--plugin", "protoc-gen-grpc=" + plugin.path)
    protoc_args.add("--grpc_out", protoc_out.path)

    protoc_args.add_all(get_include_protoc_args(proto_imports))
    protoc_args.add_all(proto_srcs, map_each = _proto_srcname)

    ctx.actions.run(
        executable = ctx.executable._protoc,
        arguments = [protoc_args],
        inputs = proto_srcs + proto_imports,
        outputs = protoc_outputs,
        tools = [plugin],
    )

    return DefaultInfo(files = depset(rule_outputs))

cc_grpc_codegen = rule(
    implementation = _cc_grpc_codegen,
    attrs = {
        "proto": attr.label(
            mandatory = True,
            allow_single_file = True,
            providers = [ProtoInfo],
        ),
        "_protoc_gen_grpc": attr.label(
            default = Label("@com_github_grpc_grpc//src/compiler:grpc_cpp_plugin"),
            executable = True,
            cfg = "host",
        ),
        "_protoc": attr.label(
            default = Label("//external:protocol_compiler"),
            executable = True,
            cfg = "host",
        ),
    },
)
