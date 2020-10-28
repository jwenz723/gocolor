# Argo CD

This directory contains various [Application CRD](https://argoproj.github.io/argo-cd/operator-manual/application.yaml)
definitions for installing gocolor as an application into Argo CD.

To install all applications into argo run from the [argocd](.) directory:

```bash
kubectl apply -f .
```