# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Uninstall a Release

## Task objective

When a release is no longer needed, we can execute one command that will remove
all the k8s objects helm has created with the deployment.

In this task, we will learn how to use `helm uninstall` command to remove the release.

## Task

Remove the previously deployed release `my-deploy` and keep the history of the release by using `--keep-history` command line flag.

<details>
    <summary>Show solution</summary>

Execute the command `helm uninstall my-deploy --keep-history`

```console
# helm uninstall my-deploy --keep-history
release "my-deploy" uninstalled
```

Since we have specified to keep the history of the release, we can check the history of the deploy.

```console
# helm history my-deploy

REVISION	UPDATED                 	STATUS     	CHART                 	APP VERSION	DESCRIPTION
1       	Tue May 26 22:39:49 2020	superseded 	workshop-example-0.1.0	1.16.0     	Install complete
2       	Tue May 26 23:03:02 2020	superseded 	workshop-example-0.1.0	1.16.0     	Upgrade complete
3       	Tue May 26 23:14:07 2020	uninstalled	workshop-example-0.1.0	1.16.0     	Uninstallation complete
```

</details>
