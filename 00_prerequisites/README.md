# [Workshop](../README.md) &raquo; Workshop prerequisites

To actively participate in the workshop the following tools will be required.
For the purpose of this workshop we have pre-created a kubernetes cluster in
Google Cloud and created multiple namespaces, named by the `$CODE`. Each
participant has access to one of the namespaces, where all the kubernetes
related tasks can be performed.

To simplify access to the cluster, we have created docker images with already
baked-in credentials to access your specific namespace. The image has
preconfigured `.kubeconfig` and all the required environmental variables set.

To actively participate in the workshop the following tools will be required.

## Table of contents

1. [Github repository](#github-repository)
2. [Docker installations](#docker-installation)
3. [Access to your namespace](#access-to-your-namespace)
4. [Access to Kubernetes dashboard](#access-to-kubernetes-dashboard)

## Github repository

All the code for the hands-on part of the workshop is available within this
repository.

If you would like to have your local copy, use `git` to download it (or download
[zip](https://github.com/3fs/k8s-workshop-episode-2/archive/master.zip)
archive).

## Docker installation

For execution of the tasks docker images are already prepared, your computer
just needs to be able to run them. To achieve that, please install Docker for
your operating system.

Installation instruction can be found [here](https://docs.docker.com/get-docker/).

## Access to your namespace

Run the below command to start a docker container with credentials to access
your test namespace in a kubernetes cluster.

Substitute `${CODE}` with workshop code you received.

```bash
docker run \
    -it \
    eu.gcr.io/k8s-workshop-2/console:${CODE}
```

## Access to Kubernetes dashboard

The workshop kubernetes cluster has [kubernetes
dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
enabled. To access the dashboard:

1. open [kubernetes dashboard](https://dashboard.k8s.3fs.si)
2. login using token which can be obtained from the `console` container with
   command `login-token`
3. enter the name of your namespace into the input field on the dashboard's left
   column


## Next: [Basic objects](../01_basic_objects/README.md)
