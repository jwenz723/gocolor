# jsonnet

To compile the jsonnet source locally execute the following from the [k8s/jsonnet](.) directory:

```bash
jsonnet -J vendor ./gocolor/main.jsonnet
```

To specify top level arguments:

```bash
jsonnet -J vendor ./gocolor/main.jsonnet --tla-str port=80
```