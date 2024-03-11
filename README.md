# remote-apis

This repository contains a collection of APIs which work together to enable
large scale distributed execution and caching on source code and other inputs.
It describes how to upload inputs, request the execution, monitor for results,
and cache those results. It's overall aim is to enable large scale parallel executions
that wouldn't be feasible on a single system, while minimizing the amount of uploads
and executions needed by storing data in a content-addressable format and caching results.

### [Remote Execution API](build/bazel/remote/execution/v2/remote_execution.proto)

The Remote Execution API is an API that, at its most general, allows clients to request
execution of binaries on a remote system. It is intended primarily for use by build systems,
such as [Bazel](bazel.build), to distribute build and test actions through a worker pool,
and also provide a central cache of build results. This allows builds to execute
faster, both by reusing results already built by other clients and by allowing many
actions to be executed in parallel, in excess of the resource limits of the machine
running the build.

### [Remote Asset API](build/bazel/remote/asset/v1/remote_asset.proto)

The Remote Asset API is an API to associate Qualifiers and URIs to Digests stored in
Content Addressable Storage. It is primary intended to allow clients to use semantically
relevant identifiers, such as a git repository or tarball location, to get the corresponding
Digest. This mapping may be pushed by a client directly, or dynamically resolved and added
to CAS by the asset server when fetched by a client.

### [Remote Logstream API](build/bazel/remote/logstream/v1/remote_logstream.proto)

The Remote Logstream API is an API supporting ordered reads and writes of `LogStream`
resources. It is intented primarily for streaming the stdout and stderr of ongoing Action
executions, enabling clients to view them while the Action is executing instead of waiting
for it's completion.

## API users

There are a number of clients and services using these APIs, they are listed
below.

### Clients
These tools use the Remote Execution API to distribute builds to workers.

* [Bazel](https://bazel.build)
* [Buck2](https://github.com/facebook/buck2)
* [BuildStream](https://buildstream.build/)
* [Goma Server](https://chromium.googlesource.com/infra/goma/server/)
* [Justbuild](https://github.com/just-buildsystem/justbuild) (via `--compatible`)
* [Pants](https://www.pantsbuild.org)
* [Please](https://please.build)
* [Recc](https://gitlab.com/bloomberg/recc)
* [Reclient](https://github.com/bazelbuild/reclient)
* [Siso](https://chromium.googlesource.com/infra/infra/+/refs/heads/main/go/src/infra/build/siso/)

### Servers
These applications implement the Remote Execution API to serve build requests
from the clients above.

* [bazel-remote](https://github.com/buchgr/bazel-remote) (open source, cache only)
* [Buildbarn](https://github.com/buildbarn) (open source)
* [BuildBuddy](https://www.buildbuddy.io/) (commercial & open source)
* [Buildfarm](https://github.com/bazelbuild/bazel-buildfarm) (open source)
* [BuildGrid](https://buildgrid.build/) (open source)
* [EngFlow](https://www.engflow.com/) (commercial)
* [Flare Build Execution](https://flare.build/products/flare-build-execution) (commercial)
* [Justbuild](https://github.com/just-buildsystem/justbuild/blob/master/doc/tutorial/just-execute.org) (via `--compatible`, open source)
* [Kajiya](https://chromium.googlesource.com/infra/infra/+/refs/heads/main/go/src/infra/build/kajiya/) (open source)
* [Scoot](https://github.com/twitter/scoot) (open source)
* [Turbo Cache](https://github.com/allada/turbo-cache) (open source)

### Workers
Servers generally distribute work to a fleet of workers. While there is no
standard API for communication between the server and workers, links to the
APIs from some existing implementations are provided as a reference below.

The [Remote Worker API](https://docs.google.com/document/d/1s_AzRRD2mdyktKUj2HWBn99rMg_3tcPvdjx3MPbFidU) 
defines a generic protocol for worker and server communication, although, 
this API is considered too heavyweight for most use-cases.

*Adhering to any one of these protocols is not a requirement.*

* [Buildfarm Operation Queues](https://bazelbuild.github.io/bazel-buildfarm/docs/architecture/queues/)
  * Uses sets of queues for managing different payload requirements.
* [Buildbarn Remote Worker](https://github.com/buildbarn/bb-remote-execution/blob/master/pkg/proto/remoteworker/remoteworker.proto)
  * Uses a custom protocol for workers to connect to a scheduler and receive instructions.
* [BuildGrid Bots](https://buildgrid.build/developer/data_model.html#rwapi)
  * A server implementation of the Remote Workers API.
* [Buildbox Worker](https://gitlab.com/BuildGrid/buildbox/buildbox-worker)
  * A worker implementation of the Remote Workers API.

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
