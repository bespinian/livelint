# Livelint

A CLI linting tool for live Kubernetes deployments

Livelint is based on [the A visual guide on troubleshooting Kubernetes deployments blog post](https://learnk8s.io/troubleshooting-deployments) by Daniele Polencic.

## Installation

```shell
$ go install github.com/bespinian/livelint@latest
```

## Usage

```shell
$ livelint check my-deployment
```

## Development

We are thrilled to receive feedback, issues and pull requests from the community.

### Build

```shell
$ make build
```

### Lint Code

```shell
$ make lint
```

### Run Tests

```shell
$ make test
```

There is also a test setup that can be deployed to a Kubernetes cluster to test specific use cases. It can be created by running

```
$ kubectl apply -f testdata
```
