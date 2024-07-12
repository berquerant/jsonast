# jsonast

```
❯ jsonast --help
jsonast - build json AST

Usage:

  jsonast [flags] [FILE]

This command converts json to AST but does not convert the value type of the node.
Node values ​​are always strings.

Flags:
```

# Examples

``` shell
cat - <<EOS | jsonast | jq
{
  "hello": "world!"
}
EOS
{
  "path": ".",
  "pos": {
    "line": 1,
    "column": 1,
    "offset": 0
  },
  "object": {
    "path": ".",
    "pos": {
      "line": 1,
      "column": 1,
      "offset": 0
    },
    "pairs": [
      {
        "path": ".[\"hello\"]",
        "pos": {
          "line": 2,
          "column": 3,
          "offset": 4
        },
        "key": "hello",
        "value": {
          "path": ".[\"hello\"]",
          "pos": {
            "line": 2,
            "column": 12,
            "offset": 13
          },
          "string": "world!"
        }
      }
    ]
  }
}
```
