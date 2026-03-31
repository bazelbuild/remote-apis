#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail

# Single tag arg passed by
# https://github.com/bazel-contrib/.github/blob/master/.github/workflows/release_ruleset.yaml
TAG="$1"
VERSION="${TAG#v}"
PREFIX="remote-apis-${VERSION}"
ARCHIVE="remote-apis-$TAG.tar.gz"

# Don't archive .bazelversion symlinks, as they break Bazel as of
# bazelbuild/bazel@f942a706a39f9803e5611b1973172ece361197ac.
# See: https://github.com/bazelbuild/bazel-worker-api/issues/21
git archive --format=tar --prefix=${PREFIX}/ ${TAG} \
  ':!:**.bazelversion' \
  ':!:.bazelci/' \
  ':!:.bcr/' \
  ':!:.github/' \
  ':!:hooks/' \
  ':!:test/bcr/' \
  | gzip > $ARCHIVE
SHA=$(shasum -a 256 $ARCHIVE | awk '{print $1}')

cat << EOF
## Using Bzlmod

Paste this snippet into your \`MODULE.bazel\` file:

\`\`\`starlark
bazel_dep(name = "bazel_remote_apis", version = "${VERSION}")
\`\`\`

## Using WORKSPACE

Paste this snippet into your legacy \`WORKSPACE\` file:

\`\`\`starlark
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_remote_apis",
    sha256 = "${SHA}",
    strip_prefix = "${PREFIX}",
    url = "https://github.com/bazelbuild/remote-apis/releases/download/${TAG}/${ARCHIVE}",
)
\`\`\`
EOF
