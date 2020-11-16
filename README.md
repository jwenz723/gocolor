# Gocolor

A simple go http server that displays a single page configured by env variables. This repository contains various methods
of deploying gocolor to Kubernetes within the [k8s](k8s) directory. This repository is intended to be a playground for
discovery and should not be considered stable.

## Configuration

View the contents of `func getConfig()` within [main.go](main.go) to get a list of available env vars used for configuration.

## Run

Clone the repo then run:

`go run main.go`