# Remote Execution API

The Remote Execution API is an API that, at its most general, allows clients to
request execution of binaries on a remote system. It is intended primarily for
use by build systems, such as [Bazel](bazel.build), to distribute build and test
actions through a worker pool, and also provide a central cache of build
results. This allows builds to execute faster, both by reusing results already
built by other clients and by allowing many actions to be executed in parallel,
in excess of the resource limits of the machine running the build.

There are a number of clients and services using these APIs, they are listed
below.

### Clients
These tools use the Remote Execution API to distribute builds to workers.

* [Bazel](https://bazel.build)
* [BuildStream](https://buildstream.build/)
* [Goma Server](https://chromium.googlesource.com/infra/goma/server/)
* [Pants](https://www.pantsbuild.org)
* [Recc](https://gitlab.com/bloomberg/recc)

### Servers
These applications implement the Remote Execution API to server build requests
from the clients above. These are then distributed to workers; some of these 
workers implement the Remote Worker API.

* [Buildbarn](https://github.com/EdSchouten/bazel-buildbarn)
* [Buildfarm](https://github.com/bazelbuild/bazel-buildfarm)
* [BuildGrid](https://buildgrid.build/)
* [Remote Build Execution (Alpha)](https://blog.bazel.build/2018/10/05/remote-build-execution.html)
* [Scoot](https://github.com/twitter/scoot)

## API Community

The [Remote Execution APIs
group](https://groups.google.com/forum/#!forum/remote-execution-apis) hosts
discussions related to the APIs in this repository.

Interested parties meet monthly via VC to discuss issues related to the APIs,
and several contributors have organized occasional meetups, hack-a-thons, and
summits. Joining the email discussion group will automatically add you to the
Google Calendar invite for the monthly meeting.

## Dependencies

The APIs in this repository refer to several general-purpose APIs published by
Google in the [Google APIs
repository](https://github.com/googleapis/googleapis). You will need to refer to
packages from that repository in order to generate code using this API. If you
build the repository using the included `BUILD` files, Bazel will fetch the
protobuf compiler and googleapis automatically.

## Using the APIs

The repository contains `BUILD` files to build the protobuf library with
[Bazel](https://bazel.build/). If you wish to use them with your own project in
Bazel, you will possibly want to declare `cc_proto_library`,
`java_proto_library`, etc. rules that depend on them.

Other build systems will have to run protoc on the protobuf files, and link in
the googleapis and well-known proto types, manually.

### Go (for non-Bazel build systems)

This repository contains the generated Go code for interacting with the API via
gRPC. Get it with:

```
go get github.com/bazelbuild/remote-apis
```

Import it with, for example:

```
repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
```

## Development

Enable the git hooks to automatically generate Go proto code on commit:

```
git config core.hooksPath hooks/
```

This is a local setting, so applies only to this repository.
