# POC v8go

[docs](https://github.com/grafana/sobek?tab=readme-ov-file)


- [POC v8go](#poc-v8go)
  - [Introduction](#introduction)
  - [Build](#fast)
  - [Debugging](#debugging)


## Introduction

Loads the module and hash similar to Pepr controller

## Build
Deletes the k3d standard cluster, creates the cluster, builds the image, imports it into the cluster, runs the pod with the image.

```bash
make run-all
```

## Debugging

Get events

```bash
make events
```

Get Logs

```bash
make logs
```

Redeploy the `pod` and `configMap`


```bash
make redeploy
```

Exec into the controller pod

```bash
make k8s-exec
```

## Analysis

( If most of the work is done in javascript (for example crypto or any other heavy calculations) you are definitely better off with V8.)[https://github.com/grafana/sobek?tab=readme-ov-file#why-would-i-want-to-use-it-over-a-v8-wrapper]

- Can't seem to easily get logs
