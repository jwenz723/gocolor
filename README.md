# Gocolor

A simple go http server that displays a single page configured by env variables. This repository contains various methods
of deploying gocolor to Kubernetes within the [k8s](k8s) directory. This repository is intended to be a playground for
discovery and should not be considered stable.

## Configuration

View the contents of `func getConfig()` within [main.go](main.go) to get a list of available env vars used for configuration.

## Run

Clone the repo then run:

`go run main.go`


## Conventional Commits

This repository enforces that commits be made following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
styling. You can find a short list of the commit message `type` values that are available
[here](https://github.com/commitizen/conventional-commit-types/blob/master/index.json).
