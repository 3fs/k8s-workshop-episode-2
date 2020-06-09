# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Install an Example Chart

## Task objective

In this task, we will install pre-created `workshop-example` chart. We will use helm to deploy it and expose it
to the Internet using Ingress kubernetes object.

`workshop-example` chart files are available [here](./workshop-example/).

## Chart overview

The directory structure of the `workshop-example` chart defines several kubernetes objects, which are used as templates for `helm`. When a user want to create a release, these templates are rendered to appropriate kubernetes objects and deployed into kubernetes cluster.

```console
workshop-example
├── Chart.yaml
├── charts
├── files
│   ├── ...
├── templates
│   ├── NOTES.txt
│   ├── _helpers.tpl
│   ├── configmap.yaml
│   ├── ingress.yaml
│   ├── secret.yaml
│   ├── service-headless.yaml
│   ├── service.yaml
│   ├── statefulset.yaml
│   └── tests
│       └── test-connection.yaml
└── values.yaml
```

In this example rendered objects will create: `StatefulSet`, `Service`, `Ingress`, `ConfigMap` and `Secret` kubernetes objects.

The content of these files is configured by values in [`values.yaml`](./workshop-example/values.yaml) file.
By changing values of `values.yaml` file we can create different deployments with different set of configuration, while the core part stays the same.

## Task

Deploy `workshop-example` chart to k8s cluster with the default values to a release name `my-deploy`

<details>
    <summary>Show solution</summary>

```console
# helm install my-deploy workshop-example
NAME: my-deploy
LAST DEPLOYED: Tue May 26 22:39:49 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  https://cranky-hippo.k8s.3fs.si/
```

</details>

## Learn About Releases

You can list the deployed releases by executing the following command.

```console
# helm list
NAME     	NAMESPACE	REVISION	UPDATED                              	STATUS  	CHART                 	APP VERSION
my-deploy	default  	1       	2020-05-26 22:39:49.278465 +0200 CEST	deployed	workshop-example-0.1.0	1.16.0
```

## Next: [Inspect a Release](./02_inspect_upgrade.md)
