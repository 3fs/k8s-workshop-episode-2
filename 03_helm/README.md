# [Workshop](../README.md) &raquo; Helm basics

Helm is a tool that streamlines installing and managing Kubernetes applications.
Think of it like apt/yum/homebrew for Kubernetes.

- Helm manages releases (installations) of your charts - templates of kubernetes objects.
- Helm runs on your laptop, CI/CD, or wherever you want it to run.
- Charts are Helm packages that contain at least two things:
  - A description of the package (`Chart.yaml`)
  - One or more templates, which contain Kubernetes manifest files
- Charts can be stored on disk, or fetched from remote chart repositories
  (same as Debian or RedHat packages)

This guide covers how you can quickly get started using Helm.

## Prerequisites

The following prerequisites are required for a successful and properly secured
use of Helm.

1. A Kubernetes cluster
2. Preconfigured container used for task in previous task: [Instructions](../02_kubernetes/README.md#access-to-your-namespace-in-workshop-k8s-cluster)
3. Installed helm client (already installed in the container)

## Basic commands

This is a set of `helm` command with examples

- `helm template` - render helm chart without deploying it to k8s cluster

  ```console
  helm template -f my-values.yaml my-deploy ./workshop-example
  ```

- `helm install` - create a release `my-deploy` containing all k8s object with the configuration from `my-values.yaml`.

  ```console
  helm install -f my-values.yaml my-deploy ./workshop-example
  ```

- `helm list` - list all `helm` deployment to k8s cluster

  ```console
  helm list --all
  ```

- `helm uninstall` - remove all k8s resources previously created with `helm install`

  ```console
  helm uninstall my-deploy
  ```

- `helm history` - show all deployments of helm chart and its status

  ```
  helm history my-deploy
  ```

- `helm rollback` - rollback the release to previously deployed version

  ```
  helm rollback my-deploy `
  ```

## Introduction

For this task we have pre-created `workshop-example` chart, which sums all the previous performed tasks:

- creates `StatefulSet`, `Service` and `Ingress` objects
- can be configured to create `ConfigMap`. `Secret` and `PersistentVolumeClaims`

## Tasks

- [Install chart](./01_install_chart.md)
- [Inspect upgrade](./02_inspect_upgrade.md)
- [Rollback a release](./03_rollback_release.md)
- [Delete release](./04_delete_release.md)
