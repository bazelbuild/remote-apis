workspace(name = "bazel_remote_apis")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

_bazel_skylib_version = "1.4.0"

_bazel_skylib_sha256 = "f24ab666394232f834f74d19e2ff142b0af17466ea0c69a3f4c276ee75f6efce"

http_archive(
    name = "bazel_skylib",
    sha256 = _bazel_skylib_sha256,
    urls = ["https://github.com/bazelbuild/bazel-skylib/releases/download/{0}/bazel-skylib-{0}.tar.gz".format(_bazel_skylib_version)],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

# This must be above the download of gRPC (in C++ section) and
# rules_gapic_repositories because both depend on rules_go and we need to manage
# our version of rules_go explicitly rather than depend on the version those
# bring in transitively.
_io_bazel_rules_go_version = "0.34.0"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "16e9fca53ed6bd4ff4ad76facc9b7b651a89db1689a2877d6fd7b82aa824e366",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(_io_bazel_rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(_io_bazel_rules_go_version),
    ],
)

# Gazelle dependency version should match gazelle dependency expected by gRPC
_bazel_gazelle_version = "0.24.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(_bazel_gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(_bazel_gazelle_version),
    ],
)

# Explicitly declaring Protobuf version, while Protobuf dependency is already
# instantiated in grpc_deps().
http_archive(
    name = "com_google_protobuf",
    sha256 = "0b0395d34e000f1229679e10d984ed7913078f3dd7f26cf0476467f5e65716f4",
    strip_prefix = "protobuf-23.2",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v23.2.tar.gz"],
)

_grpc_version = "1.55.1"

_grpc_sha256 = "17c0685da231917a7b3be2671a7b13b550a85fdda5e475313264c5f51c4da3f8"

http_archive(
    name = "com_github_grpc_grpc",
    sha256 = _grpc_sha256,
    strip_prefix = "grpc-%s" % _grpc_version,
    urls = ["https://github.com/grpc/grpc/archive/v%s.zip" % _grpc_version],
)

# Pull in all gRPC dependencies.
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()


# More gRPC dependencies. grpc_extra_deps does not work out of the box.
load("@build_bazel_rules_apple//apple:repositories.bzl", "apple_rules_dependencies")
load("@build_bazel_apple_support//lib:repositories.bzl", "apple_support_dependencies")

apple_rules_dependencies()

apple_support_dependencies()

bind(
    name = "grpc_cpp_plugin",
    actual = "@com_github_grpc_grpc//:grpc_cpp_plugin",
)

bind(
    name = "grpc_lib",
    actual = "@com_github_grpc_grpc//:grpc++",
)

load("//:remote_apis_deps.bzl", "remote_apis_go_deps")

remote_apis_go_deps()

_rules_gapic_version = "0.28.0"

_rules_gapic_sha256 = "921aebfb7f26c110fcdafeb6b74b1eb2d114a6d52e1bc11d6c1c011cf1da2017"

http_archive(
    name = "rules_gapic",
    sha256 = _rules_gapic_sha256,
    strip_prefix = "rules_gapic-%s" % _rules_gapic_version,
    urls = ["https://github.com/googleapis/rules_gapic/archive/v%s.tar.gz" % _rules_gapic_version],
)

# Needed for the googleapis protos.
http_archive(
    name = "googleapis",
    sha256 = "1253e1e9826fa822d709221032069dbbc5079e3d26671bfcb8485b21870941d3",
    strip_prefix = "googleapis-b34897064113f06083e6d5bd217826ec2c91abe3",
    urls = ["https://github.com/googleapis/googleapis/archive/b34897064113f06083e6d5bd217826ec2c91abe3.zip"],
)

load("@googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    gapic = True,
    go = True,
    java = True,
)
