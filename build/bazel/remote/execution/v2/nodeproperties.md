# Node Property Lexicon

This lexicon defines standard property names and values that servers MAY
support in the `NodeProperty` message.

The following standard property `name`s are defined:

* `UnixMode`: This describes the access rights of the file represented by the
  node. The value of this must be a string representation of the access mode
  bits rendered as a 4-digit octal. Multiple values are not allowed.

  Example:
  ```json
  // (Directory proto)
  {
    "files": [
      {
        "name": "foo",
        "digest": {...},
        "node_properties": [
          {
            "name": "UnixMode",
            "value": "4755"
          }
        ]
      }
    ]
  }
  ```

* `MTime`: This describes the last data-modification time of a file represented
  by a node. The value of this must be a string encoding this time in the
  Timestamp format documented at https://www.ietf.org/rfc/rfc3339.txt. Multiple
  values are not allowed.

  Example:
  ```json
  // (Directory proto)
  {
    "files": [
      {
        "name": "foo",
        "digest": {...},
        "node_properties": [
          {
            "name": "MTime",
            "value": "2017-01-15T01:30:15.01Z"
          }
        ]
      }
    ]
  }
  ```
