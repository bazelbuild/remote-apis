# Remote Execution API

The Remote Execution API is an API that, at its most general, allows clients to
request execution of binaries on a remote system. It is intended primarily for
use by build systems, such as [Bazel](bazel.build), to distribute build and test
actions through a worker pool, and also provide a central cache of build
results. This allows builds to execute faster, both by reusing results already
built by other clients and by allowing many actions to be executed in parallel,
in excess of the resource limits of the machine running the build.

## Dependencies

The APIs in this repository refer to several general-purpose APIs published by
Google in the [Google APIs
repository](https://github.com/googleapis/googleapis). You will need to refer to
packages from that repository in order to generate code using this API.

## Using the APIs

We currently do not have a build system in this repository. All the APIs are
written in [gRPC](grpc.io), and can be compiled to the language of your choice
using the protobuf compiler.
