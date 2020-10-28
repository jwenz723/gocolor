# jsonnet

The gocolor jsonnet source is in a sub-directory within this directory so that the jsonnetfile.json and jsonnetfile.lock.json
files can not be contained within the same directory as main.jsonnet. This is necessary when loading this jsonnet into argocd
because the argocd jsonnet parser tries to parse the jsonnetfile.json and jsonnetfile.lock.json files as jsonnet source. This
should be fixed at some point in the near future with [this](https://github.com/argoproj/argo-cd/issues/4432) issue.

## Useful Documentation

* https://jsonnet-libs.github.io/k8s-alpha/1.18/
* https://jsonnet.org/

## Compiling Locally

To compile the jsonnet source locally execute the following from the [k8s/jsonnet](.) directory:

```bash
jsonnet -J vendor ./main/main.jsonnet
```

To specify top level arguments:

```bash
jsonnet -J vendor ./main/main.jsonnet --tla-str port=80
```