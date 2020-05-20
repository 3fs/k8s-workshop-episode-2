# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Rollback a Release

## Task objective

Because Helm tracks your previous releases you can audit a cluster's history
and rollback to a previous release with `helm rollback` command.

In this task we will learn how to rollback a release to a revision in release history.

```console
helm history my-deploy
REVISION	UPDATED                 	STATUS    	CHART                 	APP VERSION	DESCRIPTION
1       	Tue May 26 22:39:49 2020	superseded	workshop-example-0.1.0	1.16.0     	Install complete
2       	Tue May 26 23:03:02 2020	deployed  	workshop-example-0.1.0	1.16.0     	Upgrade complete
```

## Task

By using `helm rollback` rollback release `my-deploy` to revision `

To make a rollback of the release, execute the following command, which
will rollback the release to revision you have specified (in our case `1`).

<details>
    <summary>Show solution</summary>

Execute `helm rollback my-deploy 1`:

```console
# helm rollback my-deploy 1
Rollback was a success! Happy Helming!
```

Inspect the history of the release:

```console
# helm history my-deploy
REVISION	UPDATED                 	STATUS    	CHART                 	APP VERSION	DESCRIPTION
1       	Tue May 26 22:39:49 2020	superseded	workshop-example-0.1.0	1.16.0     	Install complete
2       	Tue May 26 23:03:02 2020	superseded	workshop-example-0.1.0	1.16.0     	Upgrade complete
3       	Tue May 26 23:14:07 2020	deployed  	workshop-example-0.1.0	1.16.0     	Rollback to 1
```

</details>

## Next: [Delete release](./04_delete_release.md)
