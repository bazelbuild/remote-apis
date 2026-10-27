# Platform Lexicon

This lexicon defines standard platform property names and values that servers
MAY support in the `Platform` message to facilitate interoperability. New
values can be added by submitting a PR against this repository, which requires
signing the [Google CLA](https://opensource.google/documentation/reference/cla).
If signing the Google CLA is undesirable, you may submit an issue instead.

## Platform properties

Platform properties are requirements. A server that does not accept a given
property name SHOULD reject the execution request, and a server that accepts a
property is expected to honor it. Platform properties are part of the action
digest, so changing them changes the action cache key.

The following standard property `name`s are defined:

* `OSFamily`: This describes the operating system family. Multiple values
  are not allowed and an exact match is required.

  The following standard values are defined:

    - `aix`
    - `freebsd`
    - `linux`
    - `macos`
    - `sunos`
    - `windows`

  Additional values may be defined by the server. For other POSIX systems
  the recommendation is to use the output of `uname -s` in lower case.

* `ISA`: This describes the instruction set architecture including
  instruction set extensions and versions. Multiple values are allowed. If
  multiple values are specified, they are AND-ed together: the worker is
  required to support all of the listed values.

  The following standard instruction set architecture values are defined:

    - `arm-a32` (little endian)
    - `arm-a32-be` (big endian)
    - `arm-a64` (little endian)
    - `arm-a64-be` (big endian)
    - `arm-t32` (little endian)
    - `arm-t32-be` (big endian)
    - `la64v100` (little endian)
    - `power-isa-be` (big endian)
    - `power-isa-le` (little endian)
    - `rv32g` (little endian)
    - `rv64g` (little endian)
    - `sparc-v9` (big endian)
    - `x86-32`
    - `x86-64`

  The following standard instruction set extension and version values are
  defined:

    - `arm-neon`
    - `arm-sve`
    - `arm-vfpv3`
    - `arm-vfpv4`
    - `armv6`
    - `armv7`
    - `armv8`
    - `x86-avx`
    - `x86-avx2`
    - `x86-avx-512`
    - `x86-sse4.1`
    - `x86-sse4.2`

  Additional values may be defined by the server. Vendor-neutral names are
  recommended.

  Clients SHOULD NOT request instruction set extensions or versions without
  requesting an instruction set architecture.

  Examples with multiple values:

  ```json
  // (Platform proto)
  {
    "properties": [
      {
        "name": "ISA",
        "value": "x86-64"
      },
      {
        "name": "ISA",
        "value": "x86-avx2"
      }
    ]
  }
  ```

  ```json
  // (Platform proto)
  {
    "properties": [
      {
        "name": "ISA",
        "value": "arm-a64"
      },
      {
        "name": "ISA",
        "value": "armv8"
      },
      {
        "name": "ISA",
        "value": "arm-sve"
      }
    ]
  }
  ```

  ```json
  // (Platform proto)
  {
    "properties": [
      {
        "name": "ISA",
        "value": "arm-a32"
      },
      {
        "name": "ISA",
        "value": "armv7"
      },
      {
        "name": "ISA",
        "value": "arm-vfpv4"
      }
    ]
  }
  ```

## Execution hints

Like platform properties, execution hints are arbitrary key-value string
pairs that convey hints about the execution to the remote environment.
Unlike platform properties, execution hints are advisory metadata that
is not considered part of the action key. A server MAY ignore any or all
hints, and MUST NOT reject an execution solely because it does not
recognize a hint. The interpretation of an execution hint, and its
effect on execution and caching, is implementation defined.

Two successful executions of actions that differ only by their hints
SHOULD NOT produce different results, but a server MAY fail an execution
because of its hints, for example when a hint carries an invalid API
key.

Execution hints are intended for use cases such as routing (e.g. "prefer
a worker that recently fetched a similar input root") and supplying
side-channel information that should not be part of the action key (such
as API keys, license keys or other secrets).
